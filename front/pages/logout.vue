<template>
  <div>
    <v-row>
      <v-col
        justify="center"
        align="center"
      >
        <v-progress-circular
          indeterminate
          size="320"
        >
          Logging out...
        </v-progress-circular>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
/* eslint-disable no-console */

import { Vue, Component } from 'nuxt-property-decorator'
import { sleep } from '~/utils/sleep'

  @Component({
    layout: 'blank'
  })
export default class Login extends Vue {
  mounted () {
    // Temp
    sleep(5000)
      .then(() => this.$router.push('/'))
      .catch(err => console.error(err))
  }

  head () {
    return {
      title: 'Logout'
    }
  }
}
</script>
