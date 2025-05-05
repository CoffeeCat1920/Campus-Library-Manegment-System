<script setup>
import { reactive } from 'vue'
import axios from 'axios'

const form = reactive({
  name: '',
  isbn: '',
  author: '',
  genre: ''
})

const submitForm = async () => {
  try {
    const response = await axios.post('/api/recipe', {
      name: form.name,
      isbn: form.isbn,
      author: form.author,
      genre: genre,// convert to array
    })
    alert('Book added successfully!')
  } catch (err) {
    console.error(err)
    alert('Failed to add book.')
  }
}
</script>

<template>
  <div class="container mt-5">
    <h2 class="mb-4">Add Book</h2>
    <form @submit.prevent="submitForm">
      <div class="mb-3">
        <label for="name" class="form-label">Book Name</label>
        <input type="text" v-model="form.name" class="form-control" id="name" required />
      </div>

      <div class="mb-3">
        <label for="isbn" class="form-label">ISBN Number</label>
        <input type="text" v-model="form.isbn" class="form-control" id="isbn" required />
      </div>

      <div class="mb-3">
        <label for="author" class="form-label">Author</label>
        <input type="text" v-model="form.author" class="form-control" id="author" required />
      </div>

      <div class="mb-3">
        <label for="genre" class="form-label">Genre</label>
        <input type="text" v-model="form.genre" class="form-control" id="genre" required />
        <div class="form-text">You can use a comma-separated list (e.g., Fantasy, Horror)</div>
      </div>

      <button type="submit" class="btn btn-primary">Submit</button>
    </form>
  </div>
</template>
