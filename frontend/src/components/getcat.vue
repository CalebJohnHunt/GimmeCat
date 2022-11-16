<script lang="ts" setup>
import {onBeforeMount, reactive, ref} from 'vue'
import { GetCat, GetTags } from '../../wailsjs/go/main/App'

const data = reactive({
  cat: '',
  loading: false,
  chosenTag: '',
  console: console
})

const tags = ref<string[]>()

async function getCat() {
  console.log("Called getCat")
  data.loading = true
  let v = await GetCat(data.chosenTag)
  if (v == data.cat) {
    data.loading = false
  }
  data.cat = v
  console.log("Cats gotten", data.cat)
}

onBeforeMount(async () => {
  console.log("Getting tags")
  tags.value = await GetTags()
  console.log("Got tags")
})

</script>

<template>
  <div>
    <button @click="getCat">Get cat</button>
  </div>
  <div>
    <select v-model="data.chosenTag">
      <option v-for="tag in tags" :value="tag">{{ tag }}</option>
    </select>
  </div>
  <div>
    <div v-show="data.loading">Loading...</div>
    <img :src="data.cat" @loadeddata="data.loading = false" @load="data.loading = false" @error="data.console.log('Error!')" v-show="!data.loading"/>
  </div>
</template>

<style scoped>
img {
  max-width: 600px;
  max-height: 600px;
}
</style>
