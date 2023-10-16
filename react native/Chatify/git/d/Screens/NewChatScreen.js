import {
  View,
  Text,
  StyleSheet,
  Image,
  TextInput,
  TouchableWithoutFeedback,
  Keyboard,
  KeyboardAvoidingView,
} from "react-native";
import { AntDesign } from "@expo/vector-icons";
import { DotIndicator } from "react-native-indicators";
import { useEffect, useState } from "react";
import { FlatList } from "react-native-gesture-handler";
import { useNavigation } from "@react-navigation/native";
import { searchUsers } from "../utils/actions/userAction";
import { useDispatch, useSelector } from "react-redux";

import PageContainer from "../Components/PageContainer";
import { COLORS } from "../constants";
import UserList from "../Components/UserList";
import { windowWidth } from "../utils/screenDimensions";
import { storedUsersData } from "../store/storedUsersSlice";

const NewChatScreen = () => {
  const navigation = useNavigation();

  const dispatch = useDispatch();
  const signedInUserData = useSelector((state) => state.auth.userData);

  const [isLoading, setIsLoading] = useState(false);
  const [users, setUsers] = useState();
  const [noResultsFound, setNoResultsFound] = useState(false);

  const [searchText, setSearchText] = useState("");

  useEffect(() => {
    const delayedSearch = setTimeout(async () => {
      if (searchText !== "") {
        setIsLoading(true);
        const searchedUsersResultData = await searchUsers(searchText);
        delete searchedUsersResultData[signedInUserData.userId];
        if (Object.keys(searchedUsersResultData).length === 0) {
          console.log("no users found");
          setNoResultsFound(true);
          setIsLoading(false);
        } else {
          setUsers(searchedUsersResultData);
          dispatch(storedUsersData({ newUsers: searchedUsersResultData }));
          setNoResultsFound(false);
          setIsLoading(false);
        }
      } else {
        setIsLoading(false);
        setNoResultsFound(false);
        setUsers();
      }
    }, 500);

    return () => clearTimeout(delayedSearch);
  }, [searchText]);

  const onPressLeftButtonlHandler = () => {
    navigation.goBack();
  };

  const LeftButton = (
    <AntDesign name="close" size={22} color={COLORS.mainOrangeLight2} />
  );

  const onPressHandler = (userId) => {
    navigation.navigate("ChatList", {
      selectedUserId: userId, // not a good approach to pass userData as per react native navigation docs.
      // when we have an heavy object, always use state managment, i.e Redux or Context
    });
  };

  return (
    <PageContainer
      style={{ paddingHorizontal: 0 }}
      name="New Chat"
      leftButton={LeftButton}
      onPressLeftButton={onPressLeftButtonlHandler}
      isHeaderBottomMargin={true}
    >
      <KeyboardAvoidingView
        style={styles.container}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        keyboardVerticalOffset={Platform.OS === "ios" ? 60 : 0}
      >
        <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
          <View style={styles.container}>
            <View style={styles.searchContainer}>
              <Text style={styles.toText}>To :</Text>
              <TextInput
                placeholder="Search"
                placeholderTextColor={COLORS.lightGrey}
                style={styles.textInput}
                onChangeText={(text) => setSearchText(text)}
              />
            </View>
            {!isLoading && !noResultsFound && !users && (
              <View style={styles.searchDataContainer}>
                <Image
                  style={styles.searchImage}
                  source={require("../assets/search.png")}
                />
                <Text style={styles.searchFriendsText}>Search For Friends</Text>
              </View>
            )}
            {isLoading && !noResultsFound && (
              <View style={styles.searchDataContainer}>
                <DotIndicator color={COLORS.mainOrangeLight} />
                <Text>Searching...</Text>
              </View>
            )}
            {!isLoading && noResultsFound && (
              <View style={styles.searchDataContainer}>
                <Image
                  style={styles.searchImage}
                  source={require("../assets/not-found.jpeg")}
                />
                <Text style={styles.searchFriendsText}>No Friends Found</Text>
              </View>
            )}
            {!isLoading && !noResultsFound && users && (
              <FlatList
                data={Object.keys(users)}
                renderItem={(itemData) => {
                  const userId = itemData.item;
                  const userData = users[userId];
                  return (
                    <UserList
                      userName={userData.firstName}
                      image={userData.userProfilePic}
                      banner={userData.about}
                      onPress={() => onPressHandler(userId)}
                    />
                  );
                }}
              />
            )}
          </View>
        </TouchableWithoutFeedback>
      </KeyboardAvoidingView>
    </PageContainer>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  searchContainer: {
    flexDirection: "row",
    alignItems: "center",
    height: 40,
    width: windowWidth - 30,
    paddingHorizontal: 10,
    backgroundColor: COLORS.lightGrey,
    borderRadius: 20,
  },
  toText: {
    marginRight: 10,
    fontFamily: "Raleway-ExtraBold",
    fontSize: 15,
    color: COLORS.mainOrangeLight,
  },

  textInput: {
    width: "100%",
  },
  searchImage: {
    height: 200,
    width: 200,
  },
  searchDataContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  searchFriendsText: {
    fontFamily: "Raleway-Regular",
    fontSize: 15,
  },
});

export default NewChatScreen;
