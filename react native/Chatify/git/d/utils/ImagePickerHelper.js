import * as ImagePicker from "expo-image-picker";
import { getDownloadURL, getStorage, ref, uploadBytes } from "firebase/storage";
import { Platform } from "react-native";
import uuid from "react-native-uuid";
import { firebaseApp } from "./helper/firebaseHelper";

export const launchImagePicker = async () => {
  await checkMediaPermission();
  let result = await ImagePicker.launchImageLibraryAsync({
    mediaTypes: ImagePicker.MediaTypeOptions.Images,
    allowsEditing: true,
    aspect: [1, 1],
    quality: 1,
  });

  if (!result.canceled) {
    return result.assets[0].uri;
  }
  return;
};

const checkMediaPermission = async () => {
  if (Platform.OS !== "web") {
    const permissionResult =
      await ImagePicker.requestMediaLibraryPermissionsAsync();
    console.log(permissionResult);
    if (permissionResult.granted === false) {
      return Promise.reject("We need permissino to access your photos");
    }
  }
  return Promise.resolve();
};

export const uploadImageAsync = async (uri) => {
  const app = firebaseApp;
  const blob = await new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.onload = () => {
      resolve(xhr.response);
    };
    xhr.onerror = (error) => {
      console.log(error);
      reject(new Error("Network request failed", error));
    };
    xhr.responseType = "blob";
    xhr.open("GET", uri, true);
    xhr.send(null);
  });

  const pathName = "profilePics";
  const storageRef = ref(getStorage(), `${pathName}/${uuid.v4()}`);

  await uploadBytes(storageRef, blob);

  blob.close();

  return await getDownloadURL(storageRef);
};
