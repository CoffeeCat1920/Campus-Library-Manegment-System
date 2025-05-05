<script setup>
import { watch } from "vue";

const props = defineProps({
  isbn: String,
  name: String,
  author: String,
  available: Boolean,
});
</script>

<template>
  <tr>
    <td>{{ isbn }}</td>
    <td>{{ name }}</td>
    <td>{{ author }}</td>
    <td>{{ available }}</td>
  </tr>
</template>
