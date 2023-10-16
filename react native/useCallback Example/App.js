import { StatusBar } from "expo-status-bar";
import { useCallback, useEffect, useState } from "react";
import { SafeAreaView, StyleSheet, Text, View, Button } from "react-native";
import Total from "./Total";

const App = () => {
  const [numSandWiches, setNumSandwiches] = useState(0);
  const [numWater, setNumWater] = useState(0);

  const SANDWICHE_PRIZE = 2;

  const addWater = () => {
    setNumWater(numWater + 1);
  };

  const addSandwiches = () => {
    setNumSandwiches(numSandWiches + 1);
  };

  const getTotal = useCallback(() => {
    console.log("calculating in App");
    return numSandWiches * SANDWICHE_PRIZE;
  }, [numSandWiches]);

  return (
    <View style={styles.container}>
      <SafeAreaView />
      <Text>Water: {numWater}</Text>
      <Button title="add Water" onPress={addWater} />
      <Text>Sandwiches: {numSandWiches}</Text>
      <Button title="add Sandwiches" onPress={addSandwiches} />
      {/* <Text>Total: {getTotal()}</Text> */}
      <Total getTo={getTotal} />
      <StatusBar style="auto" />
    </View>
  );
};

export default App;

const styles = StyleSheet.create({
  container: {
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
});
