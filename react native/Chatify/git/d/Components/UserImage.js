import {
  TouchableOpacity,
  Image,
  View,
  StyleSheet,
  ActivityIndicator,
} from "react-native";

import { COLORS } from "../constants";
import userimage_default from "../assets/userimage-default.png";

const UserImage = ({ uri }) => {
  const imageSource = uri ? { uri } : userimage_default;

  return (
    <TouchableOpacity
    //onPress={}
    >
      <Image style={styles.userImage} source={imageSource} />
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  userImage: {
    height: 40,
    width: 40,
    borderRadius: 30,
    borderWidth: 2,
    borderColor: COLORS.mainOrange,
  },
});

export default UserImage;
