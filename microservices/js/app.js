  var app = new Vue({
    el: '#app',
    data: {
      message: 'Hello'
    }
    // ,
    // methods: {
    //   getData() {
    //     var route = "http://go-http.minikube.local";
    //     this.$http.get(route).then(response => {

    //         console.log(response.body);
    //         app.version = response.body.Version;
        
    //       }, response => {
              
    //       });
    //   },
    //   created: function(){
    //     // this.getData();
    //   },
    // }
});