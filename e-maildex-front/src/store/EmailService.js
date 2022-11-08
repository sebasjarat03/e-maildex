import axios from "axios"

export class EmailService{
    static serverURL= 'http://localhost:3000/indexer/fuck%20you'

    static getEmail(){
        
        
        return axios.get(this.serverURL)
    }
}