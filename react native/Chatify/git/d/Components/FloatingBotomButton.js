import { View, StyleSheet, TouchableOpacity } from "react-native";
import { COLORS } from "../constants";

const FloatingBottomButton = (props) => {
  return (
    <TouchableOpacity style={styles.floatingStyle} onPress={props.onPress}>
      {props.children}
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  floatingStyle: {
    height: 40,
    width: 40,
    position: "absolute",
    bottom: 15,
    right: 15,
    borderRadius: 20,
    backgroundColor: COLORS.mainOrangeLight,
    justifyContent: "center",
    alignItems: "center",
  },
});

export default FloatingBottomButton;
