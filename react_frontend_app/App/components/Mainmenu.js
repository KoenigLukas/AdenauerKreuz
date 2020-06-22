import React from "react";
import {
  StyleSheet,
  Text,
  View,
  TextInput,
  TouchableOpacity,
  KeyboardAvoidingView,
  AsyncStorage,
} from "react-native";

export default class Mainmenu extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: this.props.navigation.state.params.username,
    };
  }

  logout = () => {
    this.props.navigation.goBack();
  };
  createNewAK = () => {
    this.props.navigation.navigate("AdenauerKreuz", { Kreuzname });
  };

  render() {
    return (
      <View>
        <Text style={styles.title}>Ardenauer Kreuz</Text>
        <Text style={styles.subTitle}>Welcome Back {this.state.username}</Text>
        <View style={styles.logout}>
          <TouchableOpacity onPress={this.createNewAK} style={styles.button}>
            <Text style={styles.buttonText}>New Ardenauer Kreuz</Text>
          </TouchableOpacity>

          <TouchableOpacity
            onPress={this.registerButtonHandler}
            style={styles.button}
          >
            <Text style={styles.buttonText}>Archive</Text>
          </TouchableOpacity>

          <TouchableOpacity style={styles.button} onPress={this.logout}>
            <Text style={styles.buttonText}>Log Out</Text>
          </TouchableOpacity>
        </View>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  title: {
    fontSize: 40,
    fontWeight: "bold",
    alignSelf: "center",
    marginTop: 25,
  },
  subTitle: {
    fontSize: 20,
    fontWeight: "bold",
    alignSelf: "center",
  },
  logout: {
    alignSelf: "stretch",
    marginTop: 100,
  },
  button: {
    width: 325,
    alignSelf: "center",
    alignItems: "center",
    padding: 20,
    backgroundColor: "#fd9644",
    marginTop: 25,
  },
  buttonText: {
    fontSize: 20,
    color: "#FFFFFF",
    fontWeight: "bold",
  },
});
