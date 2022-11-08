<template>
  <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
  <div class="modal-dialog modal-xl">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLabel"></h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <h5>Subject:</h5>
      <p class="my-2">{{ email.subject }}</p>
      <h5>From: </h5>
      <p class="my-2">{{ email.sender }}</p>
      <h5>To:</h5>
      <p class="my-2">: {{ email.recipient }}</p>
      <p class="my-4">{{ email.message }}</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
        
      </div>
    </div>
  </div>
</div>
  <div class="container mt-3">
    
    <div class="row">
      <div class="col">
        
        <p class ="h3 text-success fw-bold">E-maildex
          <i class="h2 fa fa-envelope text-success" aria-hidden="true"></i>
        </p>
        <form >
          <div class="row">
            <div class="col-md-6">
              <div class="row">
                <div class="col">
              <input type="text" class="form-control" v-model="text" placeholder="Search">
            </div>
            <div class="col">
              <button @click.prevent="(e) => searchMail()" type="button" class="btn btn-primary"> Search </button>
            </div>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
  <div class="container mt-4">
    <table id ="tableMails" class="table table-hover">
      <thead>
        <tr>
          <th scope="col">Subject</th>
          <th scope="col">From</th>
          <th scope="col">To</th>
          <th scope="col">View mail</th>
        </tr>
      </thead>
      <tbody v-for="email in mails" :key="email.id">
        
        <tr>
          <td>{{email.subject}}</td>
          <td>{{email.sender}}</td>
          <td>{{email.recipient}}</td>
          <td><button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#exampleModal" v-on:click="setEmailInformation(email)">
            <i class="far fa-eye"></i>
</button></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>

  import axios from "axios"
  export default {
    name: "IndexerView",

    data: function(){
      return{
        text:'',
        mails : [],
        email: {},
        showTable:false,
        errorMessage: null
      }
    },
    async created(){
      
    },
    
    methods:{
      async searchMail(){
        try{let response = await axios.get(`http://localhost:3000/indexer/${this.text}`).then(response => (this.mails = response.data.emails))
        console.log(response)
        this.showTable =true
        }
        catch(error){
          this.errorMessage =error
        }
        
        
     
      },

      setEmailInformation(email) {
      this.email = email;
    }
    }
    
  }
</script>
