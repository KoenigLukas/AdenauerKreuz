import React, { useState } from "react";
import {
  StyleSheet,
  Text,
  View,
  TextInput,
  TouchableOpacity,
  KeyboardAvoidingView,
  AsyncStorage,
  CheckBox,
  ScrollView,
} from "react-native";

export default class AdenauerKreuz extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      KreuzName: "Default-Kreuz",

      argumentTitle: "",
      argumentPro: false,
      argumentCon: false,
      argumentPoints: 0,

      argumentList: [
        {
          key: "1",
          argumentTitle: "Test",
          argumentPro: true,
          argumentCon: false,
          argumentPoints: 5,
        },
        {
          key: "2",
          argumentTitle: "Test2",
          argumentPro: false,
          argumentCon: true,
          argumentPoints: 9,
        },
      ],
    };
  }

  proCheckBoxChange() {
    this.setState({
      argumentPro: !this.state.argumentPro,
      argumentCon: false,
    });
  }
  conCheckBoxChange() {
    this.setState({
      argumentCon: !this.state.argumentCon,
      argumentPro: false,
    });
  }

  render() {
    return (
      <View>
        <View
          style={{
            flexDirection: "row",
            alignItems: "flex-start",
            height: 100,
          }}
        >
          <View style={styles.inputWrapArgument}>
            <Text style={styles.label}>Argument</Text>
            <TextInput
              style={styles.inputDate}
              onChangeText={(argumentTitle) => this.setState({ argumentTitle })}
            />
          </View>

          <View style={styles.inputWrap}>
            <Text style={styles.label}>Pro</Text>
            <CheckBox
              style={styles.inputDate}
              value={this.state.argumentPro}
              onChange={() => this.proCheckBoxChange()}
            />
          </View>
          <View style={styles.inputWrap}>
            <Text style={styles.label}>Con</Text>
            <CheckBox
              style={styles.inputDate}
              value={this.state.argumentCon}
              onChange={() => this.conCheckBoxChange()}
            />
          </View>

          <View style={styles.inputWrap}>
            <Text style={styles.label}>Points</Text>
            <TextInput
              style={styles.inputDate}
              keyboardType={"numeric"}
              onChangeText={(argumentPoints) =>
                this.setState({ argumentPoints })
              }
            />
          </View>

          <View style={styles.inputWrap}>
            <Text style={styles.label}>Submit</Text>
            <TouchableOpacity onPress={this.submit} style={styles.button}>
              <Text style={styles.buttonText}>Submit</Text>
            </TouchableOpacity>
          </View>
        </View>

        <ScrollView>
          {this.state.argumentList.map((item) => (
            <View key={item.key}>
              <Text>{item.argumentTitle}</Text>
            </View>
          ))}
        </ScrollView>
      </View>
    );
  }

  submit() {
    this.state.argumentList.push({
      id: 3,
      argumentTitle: this.state.argumentTitle,
      argumentPro: this.state.argumentPro,
      argumentCon: this.state.argumentCon,
      argumentPoints: this.state.argumentPoints,
    });
  }
}

const styles = StyleSheet.create({
  label: {
    flex: 1,
    fontWeight: "bold",
    textAlign: "center",
  },
  inputWrapArgument: {
    flex: 0.8,
  },
  inputWrap: {
    flex: 0.28,
    marginRight: 5,

    alignItems: "center",
    flexDirection: "column",
  },
  inputDate: {
    flex: 1,
  },
  buttonText: {
    color: "#000",
    fontWeight: "bold",
    textAlign: "center",
    alignSelf: "center",
    alignItems: "center",
    justifyContent: "center",
  },
  TextInput: {
    backgroundColor: "#8f8f8f",
  },
});
