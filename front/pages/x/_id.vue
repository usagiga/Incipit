<template>
  <div>
    <v-card class="mx-auto" min-width="320px" width="60%">
      <v-card-title>Incipit</v-card-title>
      <v-card-subtitle>Private URL shortener</v-card-subtitle>
      <v-card-text class="text-h4 text-center">
        <div>{{ srcUrl }}</div>
        <div>
          <v-icon color="grey lighten-1">
            mdi-chevron-down
          </v-icon>
        </div>
        <div><a :href="destUrl">{{ destUrl }}</a></div>
      </v-card-text>
    </v-card>

    <!-- Loading indicator -->
    <v-overlay :value="isGetQueued">
      <v-progress-circular indeterminate size="64" />
    </v-overlay>
  </div>
</template>

<script lang="ts">
/* eslint-disable camelcase */

import { Vue, Component } from 'nuxt-property-decorator'
import IncipitApi from '~/utils/incipit-api'

  @Component({
    layout: 'blank'
  })
export default class ShortID extends Vue {
    shortId: string = ''
    destUrl: string = ''
    isGetQueued: boolean = true

    get srcUrl (): string {
      if (!process.client) {
        return ''
      }

      return window.location.href
    }

    mounted () {
      // Parse parameters
      this.shortId = this.$route.params.id

      // Validate parameters
      if (this.shortId === '') {
        this.$router.push('/link')
      }

      // Get link info through API
      this.isGetQueued = true
      IncipitApi(this.$router)
        .isInstalled()
        .then(() => {
          IncipitApi(this.$router)
            .getLinkByShortID(this.shortId)
            .then((resJson: any) => {
              const resLink = resJson?.link
              this.destUrl = resLink?.url
            })
            .finally(() => {
              this.isGetQueued = false
            })
        })
    }

    head () {
      return {
        title: 'Incipit',
        titleTemplate: ''
      }
    }
}
</script>
