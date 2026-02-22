<template>
  <section class="header w-full bg-secondary text-white mb-3 px-3 py-3">
    <div class="flex w-full justify-between mx-auto max-w-6xl">
      <h1 class="text-xl">
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
    </div>
  </section>
</template>