import { child, getDatabase, off, onValue, ref, get } from "firebase/database";
import { useEffect, useState } from "react";
import { DotIndicator } from "react-native-indicators";
import { useDispatch, useSelector } from "react-redux";
import { COLORS } from "../constants";
import { setChatsData } from "../store/chatSlice";
import { setMessagesData } from "../store/messagesSlice";
import { storedUsersData } from "../store/storedUsersSlice";
import { firebaseApp } from "../utils/helper/firebaseHelper";
import MainStack from "./MainStack";

const AppStack = () => {
  const dispatch = useDispatch();
  const userData = useSelector((state) => state.auth.userData);
  const storedUserDatas = useSelector(
    (state) => state.storedUsers.storedUsersData
  );
  const [isLoadingChatData, setIsLoadingChatData] = useState(true);

  useEffect(() => {
    const app = firebaseApp;
    const dbRef = ref(getDatabase(app));
    const userChatRef = child(dbRef, `userChats/${userData.userId}`);
    const refs = [userChatRef];

    console.log("Subscribing to firebase listeners");

    onValue(userChatRef, (querySnapshot) => {
      const chatIds = querySnapshot.val()
        ? Object.values(querySnapshot.val())
        : {};

      const chatsData = {};
      let chatsFoundCount = 0;

      for (let i = 0; i < chatIds.length; i++) {
        const chatId = chatIds[i];
        const chatRef = child(dbRef, `chats/${chatId}`);
        refs.push(chatRef);

        onValue(chatRef, (chatSnapshot) => {
          chatsFoundCount++;
          const data = chatSnapshot.val();
          if (data) {
            data.key = chatSnapshot.key;

            data.users.forEach((userId) => {
              if (storedUserDatas[userId]) return;

              const userRef = child(dbRef, `users/${userId}`);
              refs.push(userRef);

              get(userRef).then((userSnapShot) => {
                const userDataWithId = {
                  [userId]: userSnapShot.val(),
                };

                dispatch(storedUsersData({ newUsers: userDataWithId }));
              });
            });

            chatsData[chatSnapshot.key] = data;
          }

          if (chatsFoundCount >= chatIds.length) {
            dispatch(setChatsData({ chatsData }));
            setIsLoadingChatData(false);
          }
        });

        const messageRef = child(dbRef, `messages/${chatId}`); // chat id

        refs.push(messageRef);

        onValue(messageRef, (messagesDataSnapshot) => {
          const messagesData = messagesDataSnapshot.val();
          dispatch(setMessagesData({ chatId, messagesData }));
        });
      }

      if (chatsFoundCount == 0) {
        setIsLoadingChatData(false);
      }
    });

    return () => {
      console.log("Unsubscribing from firebase listeners");
      refs.forEach((ref) => off(ref));
    };
  }, []);

  return (
    <>
      {isLoadingChatData && <DotIndicator color={COLORS.mainOrangeLight} />}
      {!isLoadingChatData && <MainStack />}
    </>
  );
};

export default AppStack;
