import {
  TouchableOpacity,
  Image,
  View,
  StyleSheet,
  Text,
  ActivityIndicator,
} from "react-native";
import { useState } from "react";
import { Entypo } from "@expo/vector-icons";

import { COLORS } from "../constants";
import userimage_default from "../assets/userimage-default.png";
import {
  launchImagePicker,
  uploadImageAsync,
} from "../utils/ImagePickerHelper";
import { updateUserValues } from "../utils/actions/authActions";
import { useDispatch } from "react-redux";
import { updateUserData } from "../store/authSlice";

const EditableUserImage = ({ uri, userId }) => {
  const dispatch = useDispatch();
  const imageSource = uri ? { uri } : userimage_default;
  const [image, setImage] = useState(imageSource);
  const [isloading, setIsloading] = useState(false);

  const pickAndEditImageHandler = async () => {
    try {
      const tempImage = await launchImagePicker();
      if (tempImage) {
        setIsloading(true);
        // uploading image to storage server
        const uploadedImage = await uploadImageAsync(tempImage);

        // setting image locally
        setImage({ uri: tempImage });
        // update on data server
        updateUserValues(userId, { userProfilePic: uploadedImage });
        // update on store
        const imageData = { userProfilePic: tempImage };
        dispatch(updateUserData({ updatedData: imageData }));
        setIsloading(false);
      }
    } catch (error) {
      console.log(error);
      setIsloading(false);
    }
  };

  return (
    <TouchableOpacity
      style={styles.container}
      onPress={pickAndEditImageHandler}
    >
      <View style={styles.imageContainer}>
        {isloading ? (
          <ActivityIndicator color={COLORS.mainText} />
        ) : (
          <Image style={styles.userImage} source={image} />
        )}
      </View>
      <View style={styles.editIconView}>
        <Entypo name="edit" size={18} color={COLORS.mainText2} />
      </View>
    </TouchableOpacity>
  );
};
export default EditableUserImage;

const styles = StyleSheet.create({
  container: {
    borderWidth: 2,
    borderRadius: 10,
    height: 75,
    width: 75,
    borderColor: COLORS.mainOrange,
    backgroundColor: COLORS.mainOrange,
  },
  userImage: {
    height: 70,
    width: 70,
    borderRadius: 10,
  },
  imageContainer: {
    height: 71,
    width: 71,
    justifyContent: "center",
    alignItems: "center",
  },
  editIconView: {
    height: 30,
    width: 30,
    bottom: 20,
    left: 50,
    borderRadius: 20,
    backgroundColor: COLORS.white,
    borderLeftWidth: 1,
    borderTopWidth: 1,
    borderColor: COLORS.mainOrange,
    justifyContent: "center",
    alignItems: "center",
  },
});
