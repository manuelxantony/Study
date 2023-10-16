// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyB6bdCa6DnhrBUgTPtH3bHsNy6Nc0xf-jg",
  authDomain: "chatify-dd71a.firebaseapp.com",
  databaseURL:
    "https://chatify-dd71a-default-rtdb.asia-southeast1.firebasedatabase.app",
  projectId: "chatify-dd71a",
  storageBucket: "chatify-dd71a.appspot.com",
  messagingSenderId: "518875418550",
  appId: "1:518875418550:web:ed1725077ee2621d303b17",
  measurementId: "G-2C37HQR5YJ",
};

// Initialize Firebase
export const firebaseApp = initializeApp(firebaseConfig);
