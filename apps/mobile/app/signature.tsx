import { View, Text, StyleSheet } from 'react-native';

export default function Signature() {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Firma</Text>
      <Text style={styles.subtitle}>Firma del cliente</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, justifyContent: 'center', alignItems: 'center', padding: 24 },
  title: { fontSize: 22, fontWeight: 'bold', marginBottom: 8 },
  subtitle: { color: '#666' },
});
