import { View, Text, StyleSheet } from "react-native";
import { COLORS } from "../constants";

const Bubble = ({ userName, text, isSignedInUser, isFirstChat }) => {
  return (
    <View style={styles.bubbleContainer}>
      {isFirstChat && (
        <View style={styles.firstChat}>
          <Text style={styles.firstChat}>Say, Hi</Text>
        </View>
      )}
      {!isFirstChat && (
        <View style={styles.chatBubble}>
          <Text style={styles.userNameText(isSignedInUser)}>{userName}</Text>
          <View style={styles.chatTextContainer(isSignedInUser)}>
            <Text style={styles.chatText}>{text}</Text>
          </View>
        </View>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  bubbleContainer: {},
  firstChat: {
    justifyContent: "center",
    alignItems: "center",
    color: COLORS.mainText2,
  },

  chatBubble: {
    marginVertical: 3,
  },
  userNameText: (isSignedInUser) => ({
    letterSpacing: 0.5,
    textTransform: "uppercase",
    fontFamily: "Raleway-ExtraBold",
    fontSize: 11,
    color: isSignedInUser ? COLORS.mainOrange : COLORS.lightBlue,
  }),
  chatTextContainer: (isSignedInUser) => ({
    marginVertical: 4,
    borderLeftWidth: 2,
    borderLeftColor: isSignedInUser ? COLORS.mainOrange : COLORS.lightBlue,
  }),
  chatText: {
    letterSpacing: 0.3,
    fontSize: 15,
    fontWeight: "500",
    padding: 5,
  },
});

export default Bubble;
