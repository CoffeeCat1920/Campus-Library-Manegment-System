<script setup>
import { ref, onMounted, defineProps } from "vue";
import Card from '../components/BookCard.vue'
import axios from 'axios'

const books = ref([]);

const fetchBooks = async () => {
  try {
    const response = await axios.get('/api/data/books')
    books.value = response.data
  } catch (error) {
    alert('Failed to fetch books because: ' + error.message)
  }
}
onMounted(fetchBooks);
</script>

<template>
  <div>

    <table>
      <thead>
        <tr>
          <th>ISBN</th>
          <th>Name</th>
          <th>Author</th>
          <th>Genre</th>
        </tr>
      </thead>
      <tbody>
        <div v-for="book in books">
          <Card :isbn="book.isbn" :name="book.name" :author="book.author" , :genre="genre" />
        </div>
      </tbody>
    </table>

  </div>
</template>
