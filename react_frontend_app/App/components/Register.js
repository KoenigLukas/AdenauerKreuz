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

export default class Register extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username_check: false,
      email_check: false,
      email: "",
      username: "",
      first_name: "",
      last_name: "",
      password: "",
      passwordConf: "",
    };
  }

  componentDidMount() {
    this._loadInitialState().done();
  }

  _loadInitialState = async () => {
    var value = await AsyncStorage.getItem("user");
    if (value !== null) {
      this.props.navigation.navigate("Profile");
    }
  };

  render() {
    return (
      <View style={styles.register}>
        <KeyboardAvoidingView>
          <Text style={styles.title}>New User?</Text>
          <Text style={styles.subTitle}>Register Here:</Text>
          <View style={styles.mainFormat}>
            <TextInput
              style={styles.textInput}
              placeholder="E-Mail:"
              underlineColorAndroid={"transparent"}
              onChangeText={(email) => this.setState({ email })}
            />
            <TextInput
              style={styles.textInput}
              placeholder="First Name:"
              underlineColorAndroid={"transparent"}
              onChangeText={(first_name) => this.setState({ first_name })}
            />
            <TextInput
              style={styles.textInput}
              placeholder="Last Name:"
              underlineColorAndroid={"transparent"}
              onChangeText={(last_name) => this.setState({ last_name })}
            />
            <TextInput
              style={styles.textInput}
              placeholder="User-Name:"
              underlineColorAndroid={"transparent"}
              onChangeText={(email) => this.setState({ email })}
            />
            <TextInput
              style={styles.textInput}
              placeholder="Password:"
              secureTextEntry={true}
              underlineColorAndroid={"transparent"}
              onChangeText={(password) => this.setState({ password })}
            />
            <TextInput
              style={styles.textInput}
              placeholder="Confirm Password:"
              secureTextEntry={true}
              underlineColorAndroid={"transparent"}
              onChangeText={(passwordConf) => this.setState({ passwordConf })}
            />
          </View>

          <TouchableOpacity style={styles.button} onPress={this.register}>
            <Text style={styles.buttonText}>Sign Up</Text>
          </TouchableOpacity>
        </KeyboardAvoidingView>
      </View>
    );
  }

  register = () => {
    fetch("http://localhost:8000/api/v1/check/userAvailable/{id}", {});

    fetch("http://localhost:8000/api/v1/register", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: this.state.username,
        password: this.state.password,
      }),
    })
      .then((response) => response.json())
      .then((res) => {
        if (res.success === true) {
          AsyncStorage.setItem("user", res.user);
          this.props.navigation.navigate("");
        } else {
          alert(res.message);
        }
      })
      .done();
  };
}

const styles = StyleSheet.create({
  register: {
    alignSelf: "stretch",
    marginTop: 0,
  },
  header: {
    fontSize: 48,
    color: "#000000",
    paddingBottom: 10,
    marginBottom: 40,
  },
  title: {
    fontSize: 40,
    fontWeight: "bold",
  },
  subTitle: {
    fontSize: 20,
    fontWeight: "bold",
  },
  mainFormat: {
    marginTop: 50,
    marginBottom: 75,
  },
  textInput: {
    fontSize: 24,
    alignSelf: "stretch",
    height: 40,
    color: "#000",
    borderBottomColor: "#fd9644",
    borderBottomWidth: 2,
    marginTop: 15,
  },
  button: {
    width: 325,
    alignSelf: "center",
    alignItems: "center",
    padding: 20,
    backgroundColor: "#fd9644",
    marginTop: 5,
  },
  buttonText: {
    fontSize: 20,
    color: "#FFFFFF",
    fontWeight: "bold",
  },
});
