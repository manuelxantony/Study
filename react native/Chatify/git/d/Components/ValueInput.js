import { View, Text, TextInput, StyleSheet } from "react-native";
import { Ionicons } from "@expo/vector-icons";

const ValueInput = ({
  iconName,
  inputValue,
  placeholder,
  id,
  onChangeInputText,
  ...rest
}) => {
  const onChangeText = (text) => {
    onChangeInputText(id, text);
  };

  return (
    <View style={styles.container}>
      <View style={styles.icon}>
        <Ionicons
          name={iconName}
          size={20}
          color="#7e7c7c"
          style={{ marginRight: 10, color: "#838080d7" }}
        />
      </View>
      <TextInput
        style={styles.text}
        value={inputValue}
        placeholder={placeholder}
        autoCapitalize="none"
        autoCorrect={false}
        onChangeText={onChangeText}
        {...rest}
      />
    </View>
  );
};

export default ValueInput;

const styles = StyleSheet.create({
  container: {
    flexDirection: "row",
    marginTop: 20,
    height: 40,
    width: "90%",
    borderWidth: 2,
    borderColor: "#9e9d9da9",
    borderRadius: 15,
    alignItems: "center",
    backgroundColor: "#ffffff",
  },
  icon: {
    marginHorizontal: 10,
    height: 30,
    alignItems: "center",
    justifyContent: "center",
    borderRightWidth: 1,
    borderRightColor: "#9e9d9da9",
  },
  text: {
    flex: 1,
    fontSize: 15,
  },
});
