<template>
    <div>
        {{ this.item.id }}
        {{ this.item.Text }}
        {{ this.item.status }}
    </div>
</template>
<script>
export default {
  props: {
    item: {
      type: Object,
      required: true
    }
  }
}
</script>