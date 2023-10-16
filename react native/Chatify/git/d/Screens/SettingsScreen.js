import { useState } from "react";
import { View, Text, TextInput, StyleSheet } from "react-native";
import Pressable from "react-native/Libraries/Components/Pressable/Pressable";
import { useDispatch, useSelector } from "react-redux";
import EditableInput from "../Components/EditableInput";
import PageContainer from "../Components/PageContainer";
import EditableUserImage from "../Components/EditableUserImage";
import { COLORS } from "../constants";
import { updateUserData } from "../store/authSlice";
import { updateUserValues, userLogout } from "../utils/actions/authActions";

const SettingsScreen = () => {
  const [isBlur, setIsBlur] = useState(false);
  const dispatch = useDispatch();
  const userData = useSelector((state) => state.auth.userData);

  const onMakeViewBlurHandler = (isUp) => {
    setIsBlur(isUp);
  };

  const onChangeTexthandler = (inputId, inputValue) => {
    const update = async (userId, inputValues) => {
      try {
        updateUserValues(userId, inputValues);
      } catch (error) {
        console.log(error);
      }
    };
    const updatedInputValue = { [inputId]: inputValue };
    update(userData.userId, updatedInputValue);
    dispatch(updateUserData({ updatedData: updatedInputValue }));
  };

  const logoutHandler = async () => {
    const action = userLogout();
    await dispatch(action);
  };

  return (
    <PageContainer name="Settings" style={styles.pageContainer}>
      <View style={{ opacity: isBlur ? 0.5 : 1, flex: 1 }}>
        <View style={styles.userDetailsContainer}>
          <EditableUserImage
            userId={userData.userId}
            uri={userData.userProfilePic}
          />

          <View style={styles.userNameAndMail}>
            <EditableInput
              id="firstName"
              inputText={userData.firstName}
              textSize={20}
              isModalUp={onMakeViewBlurHandler}
              onChangeText={onChangeTexthandler}
            />
            <Text style={styles.email}>{userData.email}</Text>
          </View>
        </View>
        <View style={styles.about}>
          <EditableInput
            id="about"
            headerText="About"
            inputText={
              userData.about ? userData.about : "Hi, I'm using Chatify"
            }
            textSize={15}
            isModalUp={onMakeViewBlurHandler}
            onChangeText={onChangeTexthandler}
          />
        </View>

        <View
          style={{
            position: "absolute",
            top: 0,
            right: 0,
          }}
        >
          <Pressable onPress={logoutHandler}>
            <Text style={styles.logout}>Log out</Text>
          </Pressable>
        </View>
      </View>
    </PageContainer>
  );
};

export default SettingsScreen;

const styles = StyleSheet.create({
  pageContainer: {
    backgroundColor: "white",
  },
  userDetailsContainer: {
    flexDirection: "row",
    marginTop: 20,
  },
  userImage: {
    height: 70,
    width: 70,
    borderRadius: 10,
  },
  userNameAndMail: {
    marginLeft: 15,
    justifyContent: "flex-end",
  },
  email: {
    marginTop: 10,
    color: COLORS.mainText2,
  },
  about: {
    marginTop: 20,
  },
  logout: {
    fontSize: 15,
    fontFamily: "Raleway-ExtraBold",
  },
});
