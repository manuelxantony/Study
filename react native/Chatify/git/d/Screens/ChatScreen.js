import {
  View,
  Text,
  StyleSheet,
  Pressable,
  TextInput,
  TouchableWithoutFeedback,
  Keyboard,
  KeyboardAvoidingView,
} from "react-native";
import { AntDesign, Ionicons } from "@expo/vector-icons";
import { Feather } from "@expo/vector-icons";

import { COLORS } from "../constants";
import { useCallback, useEffect, useState } from "react";
import PageContainer from "../Components/PageContainer";
import { useSelector } from "react-redux";
import Bubble from "../Components/Bubble";
import { createChat, sendTextMessage } from "../utils/actions/chatActions";
import { FlatList } from "react-native-gesture-handler";

const ChatScreen = ({ navigation, route }) => {
  const [messageText, setMessageText] = useState("");
  const [chatUsers, setChatUsers] = useState([]);
  const [chatUserName, setChatUserName] = useState("");

  const storedChats = useSelector((state) => state.chats.chatsData);
  const [chatId, setChatId] = useState(route?.params?.chatId);
  const chatData =
    (chatId && storedChats[chatId]) || route?.params?.newChatData;

  const userData = useSelector((state) => state.auth.userData);

  const storedUsersData = useSelector(
    (state) => state.storedUsers.storedUsersData
  );

  const chatMessages = useSelector((state) => {
    if (!chatId) return [];

    const messagesData = state.messages.messagesData[chatId];

    if (!messagesData) return [];

    const messagesList = [];

    for (const key in messagesData) {
      const message = messagesData[key];
      messagesList.push({
        key,
        ...message,
      });
    }
    return messagesList;
  });

  useState(() => {
    setChatUsers(chatData.users);
  }, [chatUsers]);

  const sendMessage = useCallback(async () => {
    const id = chatId;
    try {
      if (!id) {
        // first chat
        const id = await createChat(userData.userId, chatData);
        setChatId(id);
      }

      // send message
      await sendTextMessage(chatId, userData.userId, messageText);
    } catch (error) {
      console.log(error);
    }

    setMessageText("");
  }, [messageText]);

  const onPressLeftButtonlHandler = () => {
    navigation.goBack();
  };

  useEffect(() => {
    const otherUserId = chatUsers.find((uid) => uid != userData.userId);
    const otherUserData = storedUsersData[otherUserId];
    setChatUserName(`${otherUserData?.firstName} ${otherUserData?.lastName}`);
  }, [setChatUserName]);

  const LeftButtonView = (
    <Ionicons
      name="chevron-back-outline"
      size={28}
      color={COLORS.mainOrangeLight2}
    />
  );

  return (
    <PageContainer
      name={chatUserName}
      leftButton={LeftButtonView}
      onPressLeftButton={onPressLeftButtonlHandler}
      isHeaderBottomMargin={true}
    >
      <KeyboardAvoidingView
        style={styles.container}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        keyboardVerticalOffset={Platform.OS === "ios" ? 10 : 0}
      >
        <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
          <View style={styles.container}>
            <View style={styles.chatContainer}>
              {!chatId && <Bubble isFirstChat={true} />}
              {chatId && (
                <FlatList
                  data={chatMessages}
                  renderItem={({ item }) => {
                    const message = item;
                    const isOwnMessage = message.sentBy === userData.userId;
                    const userName = isOwnMessage
                      ? userData.firstName
                      : storedUsersData[message.sentBy].firstName;
                    return (
                      <Bubble
                        isSignedInUser={isOwnMessage}
                        userName={userName}
                        text={message.messageText}
                      />
                    );
                  }}
                />
              )}
            </View>
            <View style={styles.inputContainer}>
              <TextInput
                style={styles.textInput}
                value={messageText}
                onChangeText={(text) => setMessageText(text)}
              />
              {messageText === "" ? (
                <View style={styles.buttonContainer}>
                  <Pressable style={styles.buttonPlus}>
                    <AntDesign
                      name="plus"
                      size={22}
                      color={COLORS.mainOrange}
                    />
                  </Pressable>
                  <Pressable style={styles.buttonCamera}>
                    <AntDesign
                      name="camerao"
                      size={22}
                      color={COLORS.mainOrange}
                    />
                  </Pressable>
                </View>
              ) : (
                <View style={styles.buttonContainer}>
                  <Pressable style={styles.buttonSend} onPress={sendMessage}>
                    <Feather name="send" size={22} color={COLORS.mainOrange} />
                  </Pressable>
                </View>
              )}
            </View>
          </View>
        </TouchableWithoutFeedback>
      </KeyboardAvoidingView>
    </PageContainer>
  );
};

export default ChatScreen;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLORS.white,
  },
  chatContainer: {
    flex: 1,
    justifyContent: "flex-start",
  },
  inputContainer: {
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
    height: 55,
    width: "100%",
    backgroundColor: COLORS.white,
    padding: 10,
  },
  buttonContainer: {
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
  },
  textInput: {
    flex: 1,
    height: 35,
    padding: 10,
    marginRight: 5,
    backgroundColor: COLORS.lightGrey,
    borderRadius: 20,
  },
  buttonPlus: {
    paddingLeft: 2,
  },
  buttonCamera: {
    paddingLeft: 10,
  },
  buttonSend: {
    paddingHorizontal: 5,
  },
});
