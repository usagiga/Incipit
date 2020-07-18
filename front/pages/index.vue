<template>
  <div>
    <v-row>
      <v-col>
        <v-card
          class="mx-auto"
          max-width="320"
        >
          <v-card-title>
            Incipit
          </v-card-title>
          <v-card-text>
            <p>URL shortener service for private use.</p>
          </v-card-text>
          <v-card-actions>
            <v-btn text to="/login">
              Login
            </v-btn>
            <v-btn text href="https://github.com/usagiga/Incipit">
              <v-icon>
                mdi-github
              </v-icon>
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'

  @Component({
    layout: 'blank'
  })
export default class Index extends Vue {
}
</script>
