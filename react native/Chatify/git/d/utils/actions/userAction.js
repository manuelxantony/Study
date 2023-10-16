import {
  child,
  endAt,
  get,
  getDatabase,
  orderByChild,
  query,
  ref,
  startAt,
} from "firebase/database";
import { firebaseApp } from "../helper/firebaseHelper";

export const getUserData = async (userId) => {
  try {
    const app = firebaseApp;
    const dbRef = ref(getDatabase(app));
    const childRef = child(dbRef, `users/${userId}`);
    const snapshot = await get(childRef);
    return snapshot.val(); // val for fectching actual value
  } catch (error) {
    console.log(error);
    throw error;
  }
};

export const searchUsers = async (searchText) => {
  const queryText = searchText.toLowerCase();
  try {
    const app = firebaseApp;
    const dbRef = ref(getDatabase(app));
    const usersRef = child(dbRef, `users`);

    const queryRef = query(
      usersRef,
      orderByChild("fullName"),
      startAt(queryText),
      endAt(queryText + "\uf8ff")
    );

    const snapshot = await get(queryRef);
    if (snapshot.exists()) {
      return snapshot.val(); // val for fectching actual value
    }
    return {};
  } catch (error) {
    console.log(error);
    throw error;
  }
};
