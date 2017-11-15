<script>
	function remove_from_cart(id,n){
        var $obj = $(n);
        $.ajax({
            type: "POST",
            url: "/delete-from-cart",
            data: { id: id}
        }).done(function( res ) {
        if(res.Message == "success"){
            $($obj).closest("tr").hide();
        }      
        });
    }
</script>