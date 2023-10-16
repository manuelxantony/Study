import { StyleSheet } from "react-native";
import { FlatList } from "react-native-gesture-handler";
import { Entypo } from "@expo/vector-icons";

import PageContainer from "../Components/PageContainer";
import FloatingBottomButton from "../Components/FloatingBotomButton";
import { useEffect } from "react";
import { useSelector } from "react-redux";
import UserList from "../Components/UserList";

const ChatListScreen = ({ navigation, route }) => {
  const selectedUserId = route?.params?.selectedUserId;
  const signedInUserData = useSelector((state) => state.auth.userData);

  const userData = useSelector((state) => state.auth.userData);
  const chatsData = useSelector((state) => {
    console.log("loading chats data");
    return Object.values(state.chats.chatsData).sort((a, b) => {
      return new Date(b.updatedTime) - new Date(a.updatedTime);
    });
  });

  console.log(
    "from chat list screen >>>>>>>>>>>>>>>>>>>1111>>>>>>>>>",
    chatsData
  );

  const storedUsersData = useSelector(
    (state) => state.storedUsers.storedUsersData
  );

  // console.log("storedUsersData", storedUsersData);

  useEffect(() => {
    if (!selectedUserId) {
      return;
    }
    const chatUsers = [selectedUserId, signedInUserData.userId];

    const navigationProps = {
      newChatData: { users: chatUsers },
    };

    navigation.navigate("ChatScreen", navigationProps);
  }, [route?.params]);

  const onPressNewChathandler = () => {
    navigation.navigate("NewChat");
  };

  return (
    <PageContainer name="Chats">
      <FlatList
        data={chatsData}
        renderItem={({ item }) => {
          const chatId = item.key;
          const chatData = item;
          const otheruserId = chatData?.users.find(
            (id) => id !== userData.userId
          );

          const otheruser = storedUsersData[otheruserId];
          //console.log("otherUser", storedUsersData);
          return (
            <UserList
              userName={otheruser?.firstName}
              image={otheruser?.userProfilePic}
              banner={chatData?.lastMessageText || "Say, Hi"}
              dontShowIcon={true}
              onPress={() => {
                navigation.navigate("ChatScreen", {
                  chatId,
                });
              }}
            />
          );
        }}
      />
      <FloatingBottomButton onPress={onPressNewChathandler}>
        <Entypo name="new-message" size={22} color="white" />
      </FloatingBottomButton>
    </PageContainer>
  );
};

export default ChatListScreen;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
  text: {
    fontSize: 30,
    fontFamily: "Raleway-Regular",
  },
});
