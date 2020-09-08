<template>
  <div>
    <!-- Link list -->
    <v-list class="link-list">
      <!-- Header -->
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title>
            ID
          </v-list-item-title>
        </v-list-item-content>
        <v-list-item-content>
          <v-list-item-title>
            Actual URL
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <!-- If there's no item -->
      <v-list-item v-if="linkItems.length === 0">
        <v-list-item-content>
          <v-list-item-title>
            There's no item.
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <!-- If there are items -->
      <v-list-item
        v-for="(item, index) in linkItems"
        :key="item.title"
      >
        <v-list-item-content>
          <v-list-item-title v-show="!item.isEditing" v-text="item.id" />
          <v-list-item-title v-show="item.isEditing" v-text="item.id" />
        </v-list-item-content>

        <v-list-item-content>
          <v-list-item-subtitle v-show="!item.isEditing" v-text="item.actualUrl" />
          <v-text-field v-show="item.isEditing" v-model="item.editingUrl" placeholder="Actual URL" />
        </v-list-item-content>

        <v-list-item-action class="flex-row">
          <v-btn
            v-show="!item.isEditing"
            :disabled="item.isDeleteQueued"
            icon
            @click="editLink(item)"
          >
            <v-icon color="grey lighten-1">
              mdi-pencil
            </v-icon>
          </v-btn>
          <v-btn
            v-show="!item.isEditing"
            :loading="item.isDeleteQueued"
            :disabled="item.isDeleteQueued"
            icon
            @click="deleteLink(item, index)"
          >
            <v-icon color="grey lighten-1">
              mdi-delete
            </v-icon>
          </v-btn>
          <v-btn
            v-show="item.isEditing"
            :loading="item.isUpdateQueued"
            :disabled="item.isUpdateQueued"
            icon
            @click="submitEditLink(item)"
          >
            <v-icon color="grey lighten-1">
              mdi-checkbox-marked-circle-outline
            </v-icon>
          </v-btn>
          <v-btn
            v-show="item.isEditing"
            :disabled="item.isUpdateQueued"
            icon
            @click="cancelEditLink(item)"
          >
            <v-icon color="grey lighten-1">
              mdi-close-circle-outline
            </v-icon>
          </v-btn>
        </v-list-item-action>
      </v-list-item>

      <!-- Add Button -->
      <v-list-item class="justify-center">
        <v-list-item-action class="link-list-wide-container">
          <v-btn :block="true" @click.stop="showAddLinkDialog()">
            <v-icon color="grey lighten-1">
              mdi-plus
            </v-icon>
          </v-btn>
        </v-list-item-action>
      </v-list-item>
    </v-list>

    <!-- Add dialog -->
    <v-dialog
      v-model="addLinkForm.visibleAddDialog"
      :persistent="addLinkForm.isCreateQueued"
      width="640px"
    >
      <v-card>
        <v-card-title>
          Add Link
        </v-card-title>
        <v-card-text>
          <!--suppress HtmlUnknownBooleanAttribute -->
          <v-form
            ref="addLinkForm"
            v-model="addLinkForm.isValidFormValue"
            @submit.prevent
          >
            <v-text-field
              v-model="addLinkForm.addUrl"
              label="URL"
              :rules="addLinkForm.addUrlRules"
              outlined
              required
              clearable
            />
            <v-btn
              type="submit"
              :disabled="!addLinkForm.isValidFormValue"
              :loading="addLinkForm.isCreateQueued"
              @click="addLinkForm.addLink()"
            >
              Add
            </v-btn>
            <v-btn
              type="cancel"
              :disabled="addLinkForm.isCreateQueued"
              @click="addLinkForm.closeDialog()"
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
/* eslint-disable camelcase,no-console */

import { Vue, Component } from 'nuxt-property-decorator'
import { VueRouter } from 'vue-router/types/router'
import { VForm } from '~/types/v-form'
import IncipitApi from '~/utils/incipit-api'

/**
   * An item of link list
   */
class LinkItem {
    id: number
    shortId: string
    actualUrl: string
    editingUrl: string = ''
    isEditing: boolean = false
    isUpdateQueued: boolean = false
    isDeleteQueued: boolean = false

    constructor (id: number, shortId: string, actualUrl: string) {
      this.id = id
      this.shortId = shortId
      this.actualUrl = actualUrl
    }
}

/**
   * Info of add link form dialog
   */
class AddLinkForm {
    visibleAddDialog: boolean = false
    isCreateQueued: boolean = false
    isValidFormValue: boolean = false

    addUrl: string = ''
    addUrlRules: ((v: any) => boolean | string)[] = [
      v => !!v || 'URL is required'
    ]

    $refs: { [key: string]: Vue | Element | Vue[] | Element[] }
    $router: VueRouter | undefined

    get formRef (): VForm {
      return this.$refs.addLinkForm as any
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

    addLink () {
      // Validation
      this.formRef.validate()
      if (!this.isValidFormValue) {
        return
      }

      if (this.$router === undefined) {
        throw new Error('add link form\'s $router is undefined')
      }

      this.isCreateQueued = true

      // Add link through API
      IncipitApi(this.$router)
        .createLink(this.addUrl)
        .then(() => {
          // Reload this page
          location.reload()
        })
        .catch((err: any) => {
          console.error(err)
        }).finally(() => {
          this.isCreateQueued = false
        })
    }
}

  @Component
export default class LinkList extends Vue {
    addLinkForm: AddLinkForm = new AddLinkForm(this.$refs)
    linkItems: LinkItem[] = []

    isGetQueued: boolean = true

    showAddLinkDialog () {
      this.addLinkForm.openDialog()
    }

    editLink (linkItem: LinkItem) {
      linkItem.editingUrl = linkItem.actualUrl

      // Change actions visibility
      linkItem.isEditing = true
    }

    submitEditLink (linkItem: LinkItem) {
      linkItem.isUpdateQueued = true

      // Update link through API
      IncipitApi(this.$router)
        .updateLink(linkItem.id, linkItem.editingUrl)
        .then(() => {
          linkItem.actualUrl = linkItem.editingUrl
          linkItem.isEditing = false
        })
        .catch((err: any) => {
          console.error(err)
        }).finally(() => {
          linkItem.isUpdateQueued = false
        })
    }

    cancelEditLink (linkItem: LinkItem) {
      // Change actions visibility
      linkItem.isEditing = false
    }

    deleteLink (linkItem: LinkItem, index: number) {
      linkItem.isDeleteQueued = true

      // Delete link through API
      IncipitApi(this.$router)
        .deleteLink(linkItem.id)
        .then(() => {
          this.linkItems.splice(index, 1)
        })
        .catch((err: any) => {
          console.error(err)
        }).finally(() => {
          linkItem.isDeleteQueued = false
        })
    }

    // noinspection JSUnusedGlobalSymbols
    mounted () {
      // Initialize add link form's router
      this.addLinkForm.$router = this.$router

      // Load links
      this.isGetQueued = true
      IncipitApi(this.$router)
        .getLinks()
        .then((resJson: any) => {
          // Set token pair
          const resLinks = resJson?.links
          resLinks.forEach((resLink: any) => {
            return this.linkItems.push(
              new LinkItem(
                resLink?.id,
                resLink?.short_id,
                resLink?.url
              )
            )
          })
        })
        .catch((err: any) => {
          console.error(err)
        }).finally(() => {
          this.isGetQueued = false
        })
    }

    head () {
      return {
        title: 'Links'
      }
    }
}
</script>

<style lang="scss">
  .link-list {
    .link-list-wide-container {
      width: 80%;
    }
  }
</style>
