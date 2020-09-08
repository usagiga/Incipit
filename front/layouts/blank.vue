<template>
  <v-app dark>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>

    <!-- Error Message Snackbar -->
    <v-snackbar :value="hasErrorMsg" timeout="-1">
      {{ errorMsg }}

      <template v-slot:action="{ attrs }">
        <v-btn color="error" text v-bind="attrs" @click="clearErrorMsg()">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-app>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'

  @Component(
    {}
  )
export default class Blank extends Vue {
    errorMsg: string = ''

    // noinspection JSUnusedGlobalSymbols
    mounted () {
      window.addEventListener('unhandledrejection', (event) => {
        this.errorMsg = event.reason.message
      })
      window.addEventListener('error', (event) => {
        this.errorMsg = event.error.message
      })
    }

    get hasErrorMsg (): boolean {
      return this.errorMsg.length !== 0
    }

    clearErrorMsg () {
      this.errorMsg = ''
    }
}
</script>
