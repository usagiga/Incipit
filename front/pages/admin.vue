<template>
  <div>
    <!-- Admin list -->
    <v-simple-table class="admin-list">
      <!-- Header -->
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Screen Name</th>
          <th>Actions</th>
        </tr>
      </thead>

      <tbody>
        <!-- If there's no item -->
        <tr v-if="adminItems.length === 0">
          <td colspan="4">
            There's no item.
          </td>
        </tr>

        <!-- If there are items -->
        <tr
          v-for="(item, index) in adminItems"
          :key="item.title"
        >
          <td>{{ item.id }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.screenName }}</td>
          <td>
            <v-btn
              :loading="item.isDeleteQueued"
              :disabled="item.isDeleteQueued"
              icon
              @click="deleteAdmin(item, index)"
            >
              <v-icon color="grey lighten-1">
                mdi-delete
              </v-icon>
            </v-btn>
          </td>
        </tr>
      </tbody>

      <!-- Add Button -->
      <tfoot>
        <tr>
          <th colspan="4">
            <v-btn :block="true" @click.stop="showAddAdminDialog()">
              <v-icon color="grey lighten-1">
                mdi-plus
              </v-icon>
            </v-btn>
          </th>
        </tr>
      </tfoot>
    </v-simple-table>

    <!-- Add dialog -->
    <v-dialog
      v-model="addAdminForm.visibleAddDialog"
      :persistent="addAdminForm.isCreateQueued"
      width="640px"
    >
      <v-card>
        <v-card-title>
          Add Admin
        </v-card-title>
        <v-card-text>
          <!--suppress HtmlUnknownBooleanAttribute -->
          <v-form
            ref="addAdminForm"
            v-model="addAdminForm.isValidFormValue"
            @submit.prevent
          >
            <v-text-field
              v-model="addAdminForm.addName"
              label="Name"
              :rules="addAdminForm.addNameRules"
              outlined
              required
              clearable
            />
            <v-text-field
              v-model="addAdminForm.addScreenName"
              label="Screen Name"
              :rules="addAdminForm.addScreenNameRules"
              outlined
              required
              clearable
            />
            <v-text-field
              v-model="addAdminForm.addPassword"
              label="Password"
              type="password"
              :rules="addAdminForm.addPasswordRules"
              outlined
              required
              clearable
            />
            <v-btn
              type="submit"
              :disabled="!addAdminForm.isValidFormValue"
              :loading="addAdminForm.isCreateQueued"
              @click="addAdminForm.addAdmin()"
            >
              Add
            </v-btn>
            <v-btn
              type="cancel"
              :disabled="addAdminForm.isCreateQueued"
              @click="addAdminForm.closeDialog()"
            >
              Cancel
            </v-btn>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>

    <!-- Loading indicator -->
    <v-overlay :value="isGetQueued">
      <v-progress-circular indeterminate size="64" />
    </v-overlay>
  </div>
</template>

<script lang="ts">
/* eslint-disable camelcase,no-console,no-unused-expressions */

import { Vue, Component } from 'nuxt-property-decorator'
import { VueRouter } from 'vue-router/types/router'
import { VForm } from '~/types/v-form'
import IncipitApi from '~/utils/incipit-api'

/**
   * An item of admin list
   */
class AdminItem {
    id: number
    name: string
    screenName: string
    isDeleteQueued: boolean = false

    constructor (id: number, name: string, screenName: string) {
      this.id = id
      this.name = name
      this.screenName = screenName
    }
}

/**
   * Info of add admin form dialog
   */
class AddAdminForm {
    visibleAddDialog: boolean = false
    isCreateQueued: boolean = false
    isValidFormValue: boolean = false

    addName: string = ''
    addNameRules: ((v: any) => boolean | string)[] = [
      v => !!v || 'Name is required',
      v => v.length > 3 || 'Name must be 3- chars',
      v => v.length < 32 || 'Name must be -32 chars'
    ]

    addScreenName: string = ''
    addScreenNameRules: ((v: any) => boolean | string)[] = [
      v => !!v || 'ScreenName is required',
      v => v.length > 3 || 'ScreenName must be 3- chars',
      v => v.length < 32 || 'ScreenName must be -32 chars'
    ]

    addPassword: string = ''
    addPasswordRules: ((v: any) => boolean | string)[] = [
      v => !!v || 'Password is required',
      v => v.length > 8 || 'Password must be 8- chars',
      v => v.length < 72 || 'Password must be -72 chars'
    ]

    $refs: { [key: string]: Vue | Element | Vue[] | Element[] }
    $router: VueRouter | undefined

    get formRef (): VForm {
      return this.$refs.addAdminForm as any
    }

    constructor ($refs: { [key: string]: Vue | Element | Vue[] | Element[] }) {
      this.$refs = $refs
    }

    openDialog () {
      this.visibleAddDialog = true
    }

    closeDialog () {
      this.formRef.reset()
      this.visibleAddDialog = false
    }

    addAdmin () {
      // Validation
      this.formRef.validate()
      if (!this.isValidFormValue) {
        return
      }

      if (this.$router === undefined) {
        throw new Error('add admin form\'s $router is undefined')
      }

      this.isCreateQueued = true

      // Add admin through API
      IncipitApi(this.$router)
        .createAdmin(this.addName, this.addScreenName, this.addPassword)
        .then(() => {
          // Reload this page
          location.reload()
        })
        .finally(() => {
          this.isCreateQueued = false
        })
    }
}

  @Component
export default class AdminList extends Vue {
    addAdminForm: AddAdminForm = new AddAdminForm(this.$refs)
    adminItems: AdminItem[] = []

    isGetQueued: boolean = true

    showAddAdminDialog () {
      this.addAdminForm.openDialog()
    }

    deleteAdmin (adminItem: AdminItem, index: number) {
      adminItem.isDeleteQueued = true

      // Delete admin through API
      IncipitApi(this.$router)
        .deleteAdmin(adminItem.id)
        .then(() => {
          this.adminItems.splice(index, 1)
        })
        .finally(() => {
          adminItem.isDeleteQueued = false
        })
    }

    // noinspection JSUnusedGlobalSymbols
    mounted () {
      // Initialize add admin form's router
      this.addAdminForm.$router = this.$router

      // Load admins
      this.isGetQueued = true
      IncipitApi(this.$router)
        .getAdmins()
        .then((resJson: any) => {
          const resAdmins = resJson?.admin_users
          resAdmins.forEach((resAdmin: any) => {
            return this.adminItems.push(
              new AdminItem(
                resAdmin?.id,
                resAdmin?.name,
                resAdmin?.screen_name
              )
            )
          })
        })
        .finally(() => {
          this.isGetQueued = false
        })
    }

    head () {
      return {
        title: 'Admins'
      }
    }
}
</script>

<style lang="scss">

</style>
