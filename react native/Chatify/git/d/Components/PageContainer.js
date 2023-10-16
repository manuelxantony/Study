import { View, Text, Pressable, StyleSheet } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { COLORS } from "../constants";
import { windowWidth } from "../utils/screenDimensions";

const PageContainer = (props) => {
  const insets = useSafeAreaInsets();
  return (
    <View style={[styles.container(insets), props.style]}>
      <View style={styles.header(props.isHeaderBottomMargin)}>
        <Pressable style={styles.leftButton} onPress={props.onPressLeftButton}>
          {props.leftButton}
        </Pressable>
        <View style={styles.nameContainer}>
          <Text style={styles.name} numberOfLines={1}>
            {props.name}
          </Text>
        </View>
        <Pressable
          style={styles.rightButton}
          onPress={props.onPressRightButton}
        >
          <Text style={styles.sideButtonText}>{props.rightButtonTitle}</Text>
        </Pressable>
      </View>
      {props.isHeaderBottomMargin && (
        <View style={styles.marginBottomHeader}></View>
      )}
      {props.children}
    </View>
  );
};

export default PageContainer;

const styles = StyleSheet.create({
  container: (insets) => ({
    flex: 1,
    backgroundColor: "#fff",
    paddingTop: insets.top,
    paddingHorizontal: 10,
    paddingBottom: insets.bottom - 20,
  }),
  header: (isHeaderBottomMargin) => ({
    flexDirection: "row",
    width: windowWidth - 20, // padding minus
    alignItems: "center",
    justifyContent: "space-between",

    // marginVertical: isHeaderBottomMargin ? 10 : null,
    // borderBottomWidth: isHeaderBottomMargin ? 0.5 : null,
    // opacity: isHeaderBottomMargin ? 0.3 : 1,
    // borderBottomColor: isHeaderBottomMargin ? COLORS.mainOrangeLight :,
  }),
  nameContainer: {
    alignItems: "center",
    width: (windowWidth - 30) / 3, // minus padding minus margin 20 -10
  },
  name: {
    fontSize: 17,
    fontFamily: "Raleway-ExtraBold",
    color: COLORS.mainText,
  },
  leftButton: {
    marginLeft: 5,
    width: (windowWidth - 30) / 3, // minus padding minus margin 20 -10
    justifyContent: "center",
  },
  rightButton: {
    marginRight: 5,
    width: (windowWidth - 30) / 3, // minus padding minus margin 20 -10
    alignItems: "flex-end",
  },
  sideButtonText: {
    fontFamily: "Raleway-ExtraBold",
    color: COLORS.mainText,
    fontSize: 15,
  },
  marginBottomHeader: {
    marginVertical: 10,
    borderBottomWidth: 0.5,
    opacity: 0.3,
    borderBottomColor: COLORS.mainOrangeLight,
  },
});
