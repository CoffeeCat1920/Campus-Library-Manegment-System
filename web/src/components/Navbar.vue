<script setup>
import { useRouter } from "vue-router";
import { onMounted } from "vue";

const router = useRouter();
</script>

<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light shadow-sm">
    <div class="container-fluid">
      <!-- Logo -->
      <router-link class="navbar-brand fw-bold" to="/">
        Library Admin Panel
      </router-link>

      <!-- Nav links -->
      <div class="d-flex">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0 flex-row">
          <li class="nav-item me-3">
            <router-link class="nav-link" to="/">Home</router-link>
          </li>
          <li class="nav-item me-3">
            <router-link class="nav-link" to="/manageBooks">Manage Books</router-link>
          </li>
          <li class="nav-item me-3">
            <router-link class="nav-link" to="/addBooks">Add Books</router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script setup>
</script>

<style scoped>
.nav-link {
  font-weight: 500;
}
</style>
