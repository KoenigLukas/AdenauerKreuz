import React from "react";
import { StyleSheet, Text, View } from "react-native";
import Registration from "./App/components/Register";
import Navigator from "./App/routes/homeStack";

export default class App extends React.Component {
  render() {
    return (
      <View style={styles.container}>
        <Navigator />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    backgroundColor: "#f7b731",
    paddingLeft: 15,
    paddingRight: 15,
    paddingBottom: 15,
    paddingTop: 22,
  },
});
