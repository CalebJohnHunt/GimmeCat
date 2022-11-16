<script lang="ts" setup>
import {onBeforeMount, reactive, ref} from 'vue'
import { GetCat, GetTags } from '../../wailsjs/go/main/App'

const data = reactive({
  cat: '',
  loading: false,
  chosenTag: '',
  console: console,
  error: ''
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

function setError(err: string) {
  data.loading = false
  data.cat = ''
  data.error = err
  setTimeout(() => {
    data.error = ''
  }, 2000)
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
    <img :src="data.cat"
        @loadeddata="data.loading = false"
        @load="data.loading = false"
        @error="setError('Image couldn\'t load. Probably a bad url :(')"
        v-show="!data.loading && data.error == ''"/>
    <h2 class="error" v-show="data.error != ''">{{ data.error }}</h2>
  </div>
</template>

<style scoped>
img {
  max-width: 600px;
  max-height: 600px;
}

.error {
  color: red
}
</style>
