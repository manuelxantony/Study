import { View, Text, StyleSheet, Pressable } from "react-native";
import { MaterialCommunityIcons } from "@expo/vector-icons";

import { COLORS } from "../constants";

const SocialButton = ({ iconName, inputText, onPress, textColor, ...rest }) => {
  return (
    <Pressable style={styles.container} {...rest} onPress={onPress}>
      <View style={styles.icon}>
        <MaterialCommunityIcons name={iconName} size={28} color={textColor} />
      </View>
      <View style={styles.textContainer}>
        <Text style={styles.text(textColor)}>{inputText}</Text>
      </View>
    </Pressable>
  );
};

export default SocialButton;

const styles = StyleSheet.create({
  container: {
    flexDirection: "row",
    marginTop: 20,
    height: 50,
    width: "90%",
    borderRadius: 8,
    backgroundColor: COLORS.white,
  },
  icon: {
    marginHorizontal: 10,
    height: 50,
    alignItems: "center",
    justifyContent: "center",
  },
  textContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    marginRight: 30,
  },
  text: (textColor) => ({
    color: textColor,
    fontSize: 17,
    fontWeight: "500",
  }),
});
