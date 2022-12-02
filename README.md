# IAB GPP Parsing Lib

This is a golang library for parsing the gpp privacy string

See: https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform

This should still be considered alpha software, the header parsing needs more work to handle edge cases.

### Example

```
 	content,err := gpp.Decode("BDACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN")
    
    
    if(err==nil){ 
    
       if(content.HasSection("uspv1") {
         content.UspV1.Version //1
         content.UspV1.Notice //1
         content.UspV1.LSPA //2 
       }
       
       if(content.HasSection("tcfeuv2") {
         content.TcfeuV2.CoreString.Version //2
         content.TcfeuV2.CoreString.CmpId //31
       }
       
    
    }

```