function searchAuthors() {

    var data = JSON.stringify({
        "id":$("#author_name").val()
    });  
    
    $.ajax({ 

         type: 'POST',   
         url: "http://localhost:8081/search-authors",
         data: data, 
         contentType: "application/json; charset=utf-8",
         dataType   : "json",
         success : function(text) {
           
            var authorResult = text;
            if(authorResult.errormessage) {
            
                $("#author_id_result").html("");
                $("#author_name_result").html("");
                $("#author_best_book_result").html("");
                $("#author_alias_name_result").html("");
                    
                return false;
            }
            $("#author_id_result").html(authorResult.id);
            $("#author_name_result").html(authorResult.name);
            $("#author_best_book_result").html(authorResult.best_book);
            $("#author_alias_name_result").html(authorResult.alias_name);
            
         }, error: function(xhr, status, error) {
             
             var errs = "searchAuthors: "+JSON.stringify(status);
             console.log(errs);

         }

    });
    

}

function searchBooks() {
    
        var data = JSON.stringify({
            "id":$("#book_name").val()
        });  
        
        $.ajax({ 
    
             type: 'POST',   
             url: "http://localhost:8081/search-books",
             data: data, 
             contentType: "application/json; charset=utf-8",
             dataType   : "json",
             success : function(text) {
               
                var bookResult = text;
                if(bookResult.errormessage) {
                
                    $("#book_id_result").html("");
                    $("#book_name_result").html("");
                    $("#book_author_name_result").html("");
                    $("#publish_date_result").html("");
                        
                    return false;
                }
                
                $("#book_id_result").html(bookResult.id);
                $("#book_name_result").html(bookResult.name);
                $("#book_author_name_result").html(bookResult.author_name);
                $("#publish_date_result").html(bookResult.publish_date);
                
             }, error: function(xhr, status, error) {
                 
                 var errs = "searchAuthors: "+JSON.stringify(status);
                 console.log(errs);
    
             }
    
        });
        
    
}