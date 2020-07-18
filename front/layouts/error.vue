<template>
  <div>
    <v-row>
      <v-col>
        <h1 v-if="error.statusCode === 404">
          {{ pageNotFound }}
        </h1>
        <h1 v-else>
          {{ otherError }}
        </h1>
        <NuxtLink to="/">
          Home page
        </NuxtLink>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'nuxt-property-decorator'

  @Component({
    layout: 'blank'
  })
export default class Blank extends Vue {
    @Prop({ default: null })
    readonly error!: any

    pageNotFound = '404 Not Found'
    otherError = 'An error occurred'

    // noinspection JSUnusedGlobalSymbols
    head () {
      const title =
        this.error.statusCode === 404 ? this.pageNotFound : this.otherError
      return {
        title
      }
    }
}
</script>
