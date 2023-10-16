import { useState } from "react";
import {
  View,
  Modal,
  Pressable,
  TextInput,
  Text,
  StyleSheet,
} from "react-native";
import { COLORS } from "../constants";
import { windowWidth } from "../utils/screenDimensions";
import { validateInput } from "../utils/actions/inputValidation";
import Button from "./Button";

const SettingsModal = ({
  id,
  inputText,
  setModalVisible,
  visible,
  onChangeText,
}) => {
  const [value, setValue] = useState(inputText);
  const [error, setError] = useState("");

  const onChangeTextHandler = (text) => {
    setValue(text);
    const result = validateInput(id, text);
    setError(result);
  };

  const okHandler = () => {
    if (error === undefined) {
      console.log("okHandler no error, calling on change text");
      onChangeText(value);
    }
    setModalVisible(false);
    //setValue(inputText); // check the cache problem
  };

  return (
    <View style={styles.modalContainer}>
      <Modal visible={visible} animationType="slide" transparent={true}>
        <View style={styles.modalContainer}>
          <View style={styles.editContainer}>
            <View style={styles.editBannerContainer}>
              <Text style={styles.editBannerText}>Edit</Text>
            </View>
            <View style={styles.editInputContainer}>
              <TextInput value={value} onChangeText={onChangeTextHandler} />
            </View>
            <Text>{error}</Text>
            <Button name="Ok" onPress={okHandler} />
          </View>
        </View>
      </Modal>
    </View>
  );
};

export default SettingsModal;

const styles = StyleSheet.create({
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
