<template>
  <section class="header w-full bg-secondary text-white px-3 py-3" :class="{'mb-3': !open}">
    <div class="flex w-full justify-between mx-auto max-w-6xl">
      <h1 class="text-xl flex items-center">
        <a href="/" class="font-bold">PocketCms.io</a>
      </h1>
      <nav class="desktop hidden md:inline">
        <ul class="flex text-sm mt-2 list-disc space-x-10">
          <li class=""><a href="/">Startseite</a></li>
          <li class="">
            <NuxtLink to="/de/demos">
              Demos
            </NuxtLink>
          </li>
          <li class=""><a href="/">Docs</a></li>
        </ul>
      </nav>
      <nav class="md:hidden text-white">
        <button v-if="!open" @click="open = !open" class="btn btn-sm btn-primary">
          <FontAwesomeIcon :icon="['fas', 'bars']" />
        </button>
        <button v-else @click="open = !open" class="btn btn-sm btn-primary">
          <FontAwesomeIcon :icon="['fas', 'times']" />
        </button>
      </nav>
    </div>
  </section>

  <section v-if="open" :class="{'mb-3': open}" class="bg-white px-3 py-3">
    <section class="grid grid-cols-6 gap-3">
      <div class="col-span-3">
        <a href="/" class="btn btn-block btn-sm btn-secondary">Startseite</a>
      </div>
      <div class="col-span-3">
        <a href="/" class="btn btn-block btn-sm btn-secondary">Startseite</a>
      </div>
      <div class="col-span-3">
        <a href="/demos" class="btn btn-block btn-sm btn-secondary">Demos</a>
      </div>
      <div class="col-span-3">
        <a href="/impressum" class="btn btn-block btn-sm btn-secondary">Impressum</a>
      </div>
    </section>
  </section>
</template>

<script setup lang="ts">
import { useLocalStorage } from "@vueuse/core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

const open = useLocalStorage('mobile-nav', false, {});

onMounted(()=>{
  open.value = false;
})
</script>