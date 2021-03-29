<template>
    <div class="background">
        <h2 class="UserContacts">My Contacts</h2>
        <img class="dots" src="../assets/more_horiz-24px.svg" alt="dotsPicture">
        <img class="plus" src="../assets/add-24px.svg" alt="plusPicture">
        <div>
            <input class="search" type="text" v-model="search" placeholder="Search in Contacts"/>
            <img src="../assets/../assets/search-24px.svg" alt="">
        </div>
        <ol class="list">
            <el-row>
                <li>Merilin Pitsina</li>
                <el-checkbox v-model="checked"></el-checkbox>
            </el-row>
            <el-row>
                <li>Nicole Angelova</li>
                <el-checkbox v-model="checked"></el-checkbox>
            </el-row>
            <el-row>
                <li>Velko Bonev</li>
                <el-checkbox v-model="checked"></el-checkbox>
            </el-row>
            <el-row>
                <li>Kostadin Kolchev</li>
                <el-checkbox v-model="checked"></el-checkbox>
            </el-row>
        </ol>
    </div>
</template>

<script>
export default {
  name: 'Contacts',
  data: () => ({
    search: '',
  }),
};
</script>

<style>
.background{
     height: 100vh;
     background-image: url('../assets/background.svg');
     background-repeat: no-repeat;
     background-size: 100%;
 }

.dots{
     position: fixed;
     top: 0px;
     left: 0px;
 }

 .plus{
     position: absolute;
     top: 0px;
     right: 0px;
 }

 .search{
     color: #008080;
     padding: 2vw;
     background: url(images/comment-author.gif) no-repeat scroll 7px 7px;
     padding-left:30px;
 }

 .list{
    display:table;
    margin: 0 auto;
    list-style-type: none;
 }

 li{
    float:left;
    width:100%;
    border-bottom: 1px solid black;
 }
 el-row{
     margin:10vw;
 }

</style>
