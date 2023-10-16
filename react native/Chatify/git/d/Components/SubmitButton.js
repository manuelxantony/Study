import { Text, StyleSheet, Pressable } from "react-native";
import { COLORS } from "../constants";

const SubmitButton = ({ submitLabel, onPress, visible, ...rest }) => {
  return (
    <Pressable
      style={({ pressed }) => [
        {
          backgroundColor: visible ? COLORS.mainOrange : COLORS.mainOrangeLight,
          opacity: pressed ? 0.9 : 1,
        },

        styles.container,
      ]}
      onPress={visible ? onPress : null}
    >
      <Text style={styles.text}>{submitLabel}</Text>
    </Pressable>
  );
};

export default SubmitButton;

const styles = StyleSheet.create({
  container: {
    marginTop: 40,
    height: 40,
    width: 100,
    borderRadius: 50,
    alignItems: "center",
    justifyContent: "center",
  },
  text: {
    fontSize: 15,
    fontWeight: "500",
    color: "#ffffff",
  },
});
