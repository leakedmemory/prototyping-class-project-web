import { View, StyleSheet, Image, TextInput } from "react-native";
import { useSession } from "../contexts/authentication";

export default function Input(props) {
  const { signIn } = useSession();
  return (
    <View style={[styles.root, props.style]}>
      <View style={styles.iconView}>{props.icon}</View>
      <TextInput style={styles.textInput} />
    </View>
  );
}

const styles = StyleSheet.create({
  root: {
    borderColor: "#828282",
    borderWidth: 1,
    padding: 10,
    borderRadius: 10,
    width: "100%",
    flexDirection: "row",
  },
  iconView: {
    width: "15%",
    justifyContent: "center",
    alignItems: "center",
  },
  textInput: {
    color: "#828282",
    lineHeight: 12,
  },
});
