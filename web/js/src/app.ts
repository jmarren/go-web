
import htmx from "htmx.org";
// import extensions
import "htmx-ext-preload"

(function() {
	htmx.defineExtension('', {
		onEvent: function(name, event): boolean {
			return true
		},

		init: function(api: any) {
		}
	})

})()






console.log("hi there!")
