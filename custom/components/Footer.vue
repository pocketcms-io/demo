<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
</script>
<template>
  <footer class="bg-secondary text-white mt-3 px-3 py-3">
    <section class="grid grid-cols-8 gap-3 mx-auto max-w-6xl text-center">
      <div class="col-span-4 md:col-span-2">
        <a href="/datenschutz" class="text-primary">
          Datenschutz
        </a>
        &
        <a href="/Impressum" class="text-primary">
          Impressum
        </a>
      </div>
      <div class="col-span-4 md:col-span-2">
        Media by
        <a href="https://unsplash.com/de/" class="text-primary">
          Unplashed
        </a>
      </div>
      <div class="col-span-4 md:col-span-2">
        Build with
        <a href="https://www.jonathan-martz.de" class="text-red-400">
          <FontAwesomeIcon :icon="['fas', 'heart']"></FontAwesomeIcon>
        </a>
        by Jonathan Martz
      </div>
      <div class="col-span-4 md:col-span-2">
        Copyright by
        <a href="https://www.jonathan-martz.de" class="text-primary">
          JMSE
        </a>
      </div>
    </section>
  </footer>
</template>