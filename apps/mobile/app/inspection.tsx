import { View, Text, StyleSheet } from 'react-native';

export default function Inspection() {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Inspección</Text>
      <Text style={styles.subtitle}>Checklist de inspección</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, justifyContent: 'center', alignItems: 'center', padding: 24 },
  title: { fontSize: 22, fontWeight: 'bold', marginBottom: 8 },
  subtitle: { color: '#666' },
});
