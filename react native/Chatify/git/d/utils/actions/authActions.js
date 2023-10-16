import { firebaseApp } from "../helper/firebaseHelper";
import {
  createUserWithEmailAndPassword,
  getAuth,
  signInWithEmailAndPassword,
} from "firebase/auth";
import { child, getDatabase, ref, set, update } from "firebase/database";
import { authenticate, logout } from "../../store/authSlice";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { getUserData } from "./userAction";
import { useDispatch } from "react-redux";

let timer;

export const signUp = (firstName, lastName, email, password) => {
  const app = firebaseApp;
  const auth = getAuth(app);
  return async (dispatch) => {
    try {
      const result = await createUserWithEmailAndPassword(
        auth,
        email,
        password
      );
      const { uid, stsTokenManager } = result.user;
      const { accessToken, expirationTime } = stsTokenManager;

      const expiryDate = new Date(expirationTime);
      const millisecondsUntilLogout = expiryDate - new Date();

      const userData = await createUser(firstName, lastName, email, uid);

      dispatch(authenticate({ token: accessToken, userData }));
      saveDataToStorage(uid, accessToken, expiryDate);

      timer = setTimeout(() => {
        dispatch(userLogout());
      }, millisecondsUntilLogout);
    } catch (error) {
      console.log(error.code);
      if (error.code === "auth/email-already-in-use") {
        throw new Error("This email is already in use");
      }
    }
  };
};

export const signIn = (email, password) => {
  const app = firebaseApp;
  const auth = getAuth(app);
  return async (dispatch) => {
    try {
      const result = await signInWithEmailAndPassword(auth, email, password);
      const { uid, stsTokenManager } = result.user;
      const { accessToken, expirationTime } = stsTokenManager;

      const expiryDate = new Date(expirationTime);
      const millisecondsUntilLogout = expiryDate - new Date();

      const userData = await getUserData(uid);

      dispatch(authenticate({ token: accessToken, userData }));
      saveDataToStorage(uid, accessToken, expiryDate);

      timer = setTimeout(() => {
        dispatch(userLogout());
      }, millisecondsUntilLogout);
    } catch (error) {
      console.log(error.code);
      if (
        error.code === "auth/user-not-found" ||
        error.code === "auth/wrong-password"
      ) {
        throw new Error("Invalid email or password");
      }
    }
  };
};

export const updateUserValues = async (userId, newUserData) => {
  const userData = await getUserData(userId);
  let updatedUserData = { ...userData, ...newUserData };

  if (newUserData.firstName || newUserData.lastName) {
    const fullName =
      `${updatedUserData.firstName} ${updatedUserData.lastName}`.toLowerCase(); // for easy searching
    updatedUserData = { ...updatedUserData, fullName };
  }

  const dbRef = ref(getDatabase());
  const userChildRef = child(dbRef, `users/${userId}`);
  await update(userChildRef, updatedUserData);
};

const createUser = async (firstName, lastName, email, userId) => {
  const fullName = `${firstName} ${lastName}`.toLowerCase(); // for easy searching
  const userData = {
    firstName,
    lastName,
    fullName,
    email,
    userId,
    signUpDate: new Date().toISOString(),
  };
  const dbRef = ref(getDatabase());
  const userChildRef = child(dbRef, `users/${userId}`);
  await set(userChildRef, userData);
  return userData;
};

const saveDataToStorage = (userId, token, expiryDate) => {
  AsyncStorage.setItem(
    "userData",
    JSON.stringify({
      userId,
      token,
      expiryDate: expiryDate.toISOString(),
    })
  );
};

export const userLogout = () => {
  return async (dispatch) => {
    AsyncStorage.clear();
    clearTimeout(timer);
    dispatch(logout());
  };
};
