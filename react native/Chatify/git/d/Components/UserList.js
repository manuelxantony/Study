import { TouchableOpacity, View, Image, Text, StyleSheet } from "react-native";
import UserImage from "./UserImage";
import { MaterialCommunityIcons } from "@expo/vector-icons";

import { COLORS } from "../constants";
import { windowWidth } from "../utils/screenDimensions";

const UserList = ({ userName, image, banner, onPress, dontShowIcon }) => {
  return (
    <TouchableOpacity style={styles.contianer} onPress={onPress}>
      <UserImage uri={image} />
      <View style={styles.dataContainer}>
        <View style={styles.userDataContainer}>
          <Text style={styles.userName}>{userName}</Text>
          <Text numberOfLines={1} style={styles.about}>
            {banner}
          </Text>
        </View>
        {!dontShowIcon && (
          <View style={styles.icon}>
            <MaterialCommunityIcons
              name="message-arrow-right-outline"
              size={26}
              color={COLORS.mainOrange}
            />
          </View>
        )}
      </View>
    </TouchableOpacity>
  );
};

export default UserList;

const styles = StyleSheet.create({
  contianer: {
    flex: 1,
    flexDirection: "row",
    height: 60,
    width: windowWidth - 30,
    borderWidth: 0.2,
    borderColor: "grey",
    borderRadius: 10,
    backgroundColor: "white",
    alignItems: "center",
    paddingHorizontal: 10,
    marginTop: 10,
    shadowOffset: {
      width: 1,
      height: 2,
    },
    shadowRadius: 1,
    shadowOpacity: 0.1,
  },
  dataContainer: {
    flex: 1,
    flexDirection: "row",
    justifyContent: "space-between",
  },
  userDataContainer: {
    marginHorizontal: 10,
  },
  userName: {
    marginBottom: 5,
    fontFamily: "Raleway-ExtraBold",
    fontSize: 16,
  },
  about: {
    fontSize: 13,
    fontFamily: "Raleway-Regular",
    color: "black",
  },
  icon: {
    justifyContent: "center",
  },
});
