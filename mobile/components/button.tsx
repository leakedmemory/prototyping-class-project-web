import { StyleSheet, Image, Text, TouchableOpacity } from "react-native";

export default function Button(props) {
  return (
    <TouchableOpacity style={styles.root} onPress={props.onPress}>
      <Text style={styles.text}>{props.text}</Text>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  root: {
    backgroundColor: "#828282",
    padding: 10,
    borderRadius: 10,
    width: "100%",
    justifyContent: "center",
    alignItems: "center",
  },
  iconView: {
    justifyContent: "center",
  },
  text: {
    color: "white",
  },
});
