import { useState } from "react";
import { View, Text, Modal, TextInput, StyleSheet } from "react-native";
import Pressable from "react-native/Libraries/Components/Pressable/Pressable";

import { COLORS } from "../constants";
import { windowWidth } from "../utils/screenDimensions";
import SettingsModal from "./SettingsModal";

const EditableInput = ({
  id,
  headerText,
  inputText,
  textSize,
  isModalUp,
  onChangeText,
}) => {
  const [editVisible, setEditVisible] = useState(false);
  const [text, setText] = useState(inputText);

  const modalVisibleHandler = (isVisible) => {
    setEditVisible(isVisible);
    isModalUp(isVisible);
  };

  const onChangeTextHandler = (text) => {
    setText(text);
    onChangeText(id, text);
  };

  return (
    <View style={styles.container}>
      <SettingsModal
        id={id}
        inputText={text}
        setModalVisible={modalVisibleHandler}
        visible={editVisible}
        onChangeText={onChangeTextHandler}
      />
      <Text style={styles.headerText}>{headerText}</Text>
      <Pressable onPress={() => modalVisibleHandler(true)}>
        <Text style={[styles.userName, { fontSize: textSize }]}>{text}</Text>
      </Pressable>
    </View>
  );
};

export default EditableInput;

const styles = StyleSheet.create({
  container: {},
  headerText: {
    fontWeight: "500",
    color: COLORS.mainText2,
    marginBottom: 7,
  },
  userName: {
    backgroundColor: "white",
    borderRadius: 15,
    fontWeight: "800",
    fontFamily: "Raleway-ExtraBold",
    color: COLORS.secondaryText,
  },
  modalContainer: {
    justifyContent: "center",
    alignItems: "center",
  },
  editContainer: {
    borderWidth: 3,
    justifyContent: "center",
    alignItems: "center",
    margin: 30,
    backgroundColor: "white",
    marginTop: 150,
    width: windowWidth - 20,
    borderRadius: 10,
  },
  editBannerContainer: {
    height: 30,
    width: "100%",
    justifyContent: "center",
    alignItems: "center",
    borderTopLeftRadius: 7,
    borderTopRightRadius: 7,
    backgroundColor: COLORS.mainOrange,
  },
  editBannerText: {
    fontFamily: "Raleway-ExtraBold",
    fontSize: 18,
  },
  editInputContainer: {
    backgroundColor: COLORS.lightGrey,
    width: "100%",
    padding: 10,
  },
  okButton: {
    justifyContent: "center",
    alignItems: "center",
    borderWidth: 2,
    height: 40,
    width: 70,
    borderRadius: 20,
    backgroundColor: COLORS.mainOrange,
    marginVertical: 10,
  },
  okText: {
    fontSize: 20,
    fontFamily: "Raleway-ExtraBold",
  },
});
