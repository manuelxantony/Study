import { Pressable, Text, StyleSheet } from "react-native";
import { COLORS } from "../constants";

const Button = ({ name, onPress }) => {
  return (
    <Pressable style={styles.okButton} onPress={onPress}>
      <Text style={styles.okText}>{name}</Text>
    </Pressable>
  );
};

export default Button;

const styles = StyleSheet.create({
  okButton: {
    justifyContent: "center",
    alignItems: "center",
    borderWidth: 2,
    // height: 40,
    // width: 70,
    paddingHorizontal: 20,
    paddingVertical: 7,
    borderRadius: 20,

    backgroundColor: COLORS.mainOrange,
    marginVertical: 10,
  },
  okText: {
    fontSize: 20,
    fontFamily: "Raleway-ExtraBold",
  },
});
