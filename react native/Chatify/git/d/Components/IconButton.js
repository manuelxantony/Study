import { Pressable, Text, StyleSheet } from "react-native";
import { COLORS } from "../constants";
import { MaterialCommunityIcons } from "@expo/vector-icons";

const IconButton = ({ name, onPress, color }) => {
  return (
    <Pressable style={styles.okButton} onPress={onPress}>
      <MaterialCommunityIcons name={name} size={30} color={color} />
    </Pressable>
  );
};

export default IconButton;

const styles = StyleSheet.create({
  okButton: {},
  okText: {
    fontSize: 20,
    fontFamily: "Raleway-ExtraBold",
  },
});
