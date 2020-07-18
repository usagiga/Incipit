<template>
  <div>
    <v-row>
      <v-col>
        <h1>
          Login
        </h1>
        <v-form
          ref="loginForm"
          v-model="isValidLogin"
        >
          <v-text-field
            v-model="userName"
            label="User Name"
            :rules="userNameRules"
            outlined
            required
            clearable
          />
          <v-text-field
            v-model="password"
            label="Password"
            type="password"
            :rules="passwordRules"
            outlined
            required
            clearable
          />
          <v-btn :disabled="!isValidLogin" @click="submit">
            Login
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
/* eslint-disable no-console */

import { Vue, Component } from 'nuxt-property-decorator'
import { VForm } from '~/types/v-form'

  @Component({
    layout: 'blank'
  })
export default class Login extends Vue {
    isValidLogin: boolean = false

    userName: string = ''
    userNameRules: ((v: any) => boolean | string)[] = [
      v => !!v || 'Name is required'
    ]

    password: string = ''
    passwordRules: ((v: any) => boolean | string)[] = [
      v => !!v || 'Password is required'
    ]

    get loginForm (): VForm {
      return this.$refs.loginForm as any
    }

    submit (): void {
      // Validation
      this.loginForm.validate()
      if (!this.isValidLogin) {
        return
      }

      // TODO : Send cred to server
      // Temp
      console.log(this.userName)
      console.log(this.password)
    }

    head () {
      return {
        title: 'Login'
      }
    }
}
</script>
