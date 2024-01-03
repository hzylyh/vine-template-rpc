/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 15:50
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 15:50
 */

package biz

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type RuleBiz struct {
	log *log.Helper
}
type Resource struct {
	Name       string
	APIVersion string
	Namespaced bool
	Kind       string
}

func (rb *RuleBiz) Add(position string, temperature int32) {
	fmt.Println(position, temperature)
}

//
//func (rb *RuleBiz) Add(info string) {
//	//var deploy appsV1.Deployment
//	fmt.Println(info)
//	resources := strings.Split(info, `---`)
//	fmt.Println(resources)
//	var (
//		decUnstructured = yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
//	)
//	apiVersions := make(map[string]Resource)
//	ctx := context.Background()
//	fmt.Println(info)
//	//clientcmd.NewDefaultClientConfig()
//	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*clientcmdapi.Config, error) {
//		kconfig, err := clientcmd.Load([]byte("apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQWVhZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJek1EWXdNakF5TWpjeU1Gb1hEVE16TURVek1EQXlNamN5TUZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTGNFCkdHbHpETk9DcWJrUFc0SzdQZG1zSDdsWUNlbmFqcmRoSDd0SnA2SzFaaXRVYUhBYnBVMG1TL1UxT2ZtM0NnOU8KQ3RGb092UDFIZ3FFSzM4Sm81VEJOS3B6OFZHalRxelJtTnExb2F0RFM5YTZFTTIxdHdHWUc4Z3lGRFJoOGtyTQpZak9OclJlSjJFNmZQa1B3SE1PN3psdnlLUU84cUlpTnFMbVhzMERxZklIQjZyUThGeml3ZERwTnMxRGl2cDE3CitNRHBHa25KQlhjbndSajdnWm1FK2w0ZUZKVFpoSDNRYXgrNTQ0MkZKZXVSMk5ISDhhVkVwMFVJaHZQZUdsajkKcnNtVzZadTVDUjVxMjVRVitFaTBlVTl5TDlCZmFGZmVNMGxsZUFabkpkOXJ6MGJLMkdLR255RXczMmdoNHlVZAplYzM5OTJmaXJCcWx6TnFzTHQwQ0F3RUFBYU5aTUZjd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZBeFN3SytmNVhVWm9meXRLbFlWVWdxanRQdVpNQlVHQTFVZEVRUU8KTUF5Q0NtdDFZbVZ5Ym1WMFpYTXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBQ01MYzE4bXdRaEhPQnpwcXBObQpnM2ZrYTF1aytqLzdoQ294eHoxMTVqcTZiaFRjNkdXaGYvMjVQME9nNkRtM1ZBN3l3TjF1UnJiLzhZc0k4b21TCklXMDdxWExiRmtONzk1K1YvQ1BlQ3ZPZUExSEtmWTNxL2NBNk9nenQrVThWbDZkZ1R4MU16NzhWNDA1WFJwVGUKK1ZITkd4U3NGMlBCMWtuWGV6eE5uVzZ4b1huanRIV1ZPRzc0OU5VdE1oTmJ6WFN6QmxmaGxrK2VUR3RmOUEwbgpOZVQ1YXIwdERoVm56eHBic2ZOV3VCTGgzVkI1TzVIWnNkZ2pwUk1jc3FBOU5XOG1Pb2MzYktNSEZGTkNuUG1jCjNMUDN2SE83LzVOMW82UkM3aXFpd3pXa2c2WjV1YzYrbUo1N0crWUxOOEg0Tk5CWGN6N2JCT1QrNElveUE4aFQKcFQ0PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    server: https://192.168.1.67:6443\n  name: kubernetes\ncontexts:\n- context:\n    cluster: kubernetes\n    user: kubernetes-admin\n  name: kubernetes-admin@kubernetes\ncurrent-context: kubernetes-admin@kubernetes\nkind: Config\npreferences: {}\nusers:\n- name: kubernetes-admin\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURJVENDQWdtZ0F3SUJBZ0lJSE15cDR0Ti8valV3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TXpBMk1ESXdNakkzTWpCYUZ3MHpNekExTXpBd01qSTNNak5hTURReApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrCmJXbHVNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXY4WmVPZ1FqcWdZbXdzWm4KNFU1TUlsYVI0S3piUUtLTHYvbmp3dWhJdWNnRFdibVhYSFo3VUhLQ2pTSlkvS0Q3THAzYU02ak5mNWJGRXNWYQpid3lRRTBURzVnZGlFb25RRTVSWWovcFd4K0F3ZUF1Q2FKTDNncHZUaHpJRkxDOXJwU2huaEFJb1NJcDhlbVdPCkd4cDFyVnBrYm5tMUxZeVdGeXVleVdVd3Rja0tBTi91MnBNc1ZMeVdKcCtDNUhWYytJalU3WkZYMy85VGxWalIKbkRrc1ZBeHFEWlYyWGxOZWFNcHVLZjh3d1BKdlJzSWFEWGdLazhoRFVzcW5IUXQvZXdlQWQ2RDNoK1FJcW91YQp5T3BScjk2ZmpacE54cEJkQ3JqdU1lbnJaNmRlTHpKTDAvdWFiaHhsV3pWWDB5V1JPREV0MGlDY0dhNHdPY3FGCkhBQ3Nmd0lEQVFBQm8xWXdWREFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUgKQXdJd0RBWURWUjBUQVFIL0JBSXdBREFmQmdOVkhTTUVHREFXZ0JRTVVzQ3ZuK1YxR2FIOHJTcFdGVklLbzdUNwptVEFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBQ3ZYbk1sVlV0SWlrWFV3YWZhSkJzL1Y4RU1zdzlZVjlqa1MxCldQS2R1WmpWR2xheHZUd1JPc3EwRjBrSVpJSTJCbkY4T3ZtaHJMdXYzcVZzY00wS0xNRmpkNzJUQjdPRU8zeE4KQVNKTVNxdGVuN3VNMjcxaG8zUkZjZFI0a1Znek03UDJPTTkzcjlNSVdYRUpKMGNYU3cwOE83Z0U2Rk9Zek5QSgpZdlJDbTU0NkhsU2Ixd015TnpmUlJHcHhhUjhSMmxFMStPUDRpV2FTaitvN3pRMUg5Y3hDWm11TzIwdVI2dEs3CnJrb2VZcmUxakdGaGpHOU1oMTJuTk15a2I3K1pLNkl1VDlaUGV6bU0wN0l4WHk3ZnllSWNTT3g3Q1prekd0WlEKQU5XL05VN1hGcWVtaXlLSDZack1TbVFzaG9ZVFFlekcrYTY5c3VJRC93M1BSQXNDMFE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBdjhaZU9nUWpxZ1ltd3NabjRVNU1JbGFSNEt6YlFLS0x2L25qd3VoSXVjZ0RXYm1YClhIWjdVSEtDalNKWS9LRDdMcDNhTTZqTmY1YkZFc1ZhYnd5UUUwVEc1Z2RpRW9uUUU1UllqL3BXeCtBd2VBdUMKYUpMM2dwdlRoeklGTEM5cnBTaG5oQUlvU0lwOGVtV09HeHAxclZwa2JubTFMWXlXRnl1ZXlXVXd0Y2tLQU4vdQoycE1zVkx5V0pwK0M1SFZjK0lqVTdaRlgzLzlUbFZqUm5Ea3NWQXhxRFpWMlhsTmVhTXB1S2Y4d3dQSnZSc0lhCkRYZ0trOGhEVXNxbkhRdC9ld2VBZDZEM2grUUlxb3VheU9wUnI5NmZqWnBOeHBCZENyanVNZW5yWjZkZUx6SkwKMC91YWJoeGxXelZYMHlXUk9ERXQwaUNjR2E0d09jcUZIQUNzZndJREFRQUJBb0lCQUEyNjBmd3RWN0JRaUczaApib1orNkl4OHppemJzMGJKWEYvK01BUkFSUlNuOFRWM1NCTGltS21GeVRyNWZvMTNxUVFKWVJuWnQveXVlcVdCClk1d1h0aWpxLzZKd1hPRWU1THJDSUNVT3N5L1VLU1F2RHNNVjRvcURSWDJrWXlKcXNyWThVa3p0S3AwbTdRcUcKUXNwRjlxN1dqZUlxN0lHNlV2U050dXlNUmhSOWxITy9KeHFkcDZwYTU0TkVIaUdGT2ZyZXEzOVRWa3hoL0ZnRgpkT3U4eVE3TloxMklDOEZnM3RCS0tTWndUcXNxY0FncFRLenBtQy9OK0VTMGZGdHdWdy9HNUtNRWgrN09ETU1yCkxxb1FNT2VtVmU1c1hVVThVT3BFdVJheloyNUFJZFVuNUlMbEs5NUdDbWk3dGhoRmVhTlQwOEtLSlAzdlFJNjEKRlJhZ2lzRUNnWUVBM2FxWURBa21FWm5tWHlSeTlGS0MzWWxTMlhFZVgrc0RxeVdiMllqclExNGczdmRpNUhLbwpscTFVK1NtcFJlMmRHZGdUNDYzNEl2VEVTZ3ZTdmhDNjRpOUs3dUxBREE5bkpzbFhlVCtMenF1LzZubURCejloClI1L0sxcTBQMWcvd3hDcWFoeVQzdWtrMWlUaUNaNzc5MnZqNW9mNTIrelFyZUFqcHh0OFJpaWtDZ1lFQTNYcUgKNU9kdGFwdEUreGRkSmlrQnpWZWE1bTBFY3RmdzdLTWxQUG10UFRtbzNGVGRtbHE0RlNWQS8vMXJQcWp2QlduTwpkNEhidWJhdnVrYnRZRzU4a0RIZ2ptd0ExelhRdGlvWm41MFMxbVRYNlBxbVBlM2o0L1FkdEQvY0prWis4ZVpoCmpZaTdETEk0T2tQNWZwZkhsSmYyVjFHMElUZzJNcnl1eHNQd0ptY0NnWUVBdktwS0oxUlNGb0VLMXNmTEg3QnEKdWV2N25CN3ZYMnhEL3hOUGh1TkdlbHhQVmZieE1NZDJQbWpnL2dFN2xjMGIvYmN2bG9XbndQamJhTHFQY0Q0RgpFaVoyZk1SNStNblRYZXZZaTlaT1JHVmlQMHVVL2tJdnBpcVhGZ2pPaGIramlSTkoydVRZQkxIeTR3dzUwZHQ1CmJUcWtIZ0RiZkF0M1EvOHlHcFlaODZFQ2dZQUtuLzQwOVFWQ3dBQW5LNC9FYk9NeFBxcW5zME1yWDBDbm1Zd0gKeU5LT3hWMWNkSlhNK2QwcTVvYWZ3VUNMMlA3ZWU1ejBEcjdEd2dmY0g5cnpiVmw4Y1dnY2JRSDlVUSthTFpyZApGV1A1OVF5R21MK3c2T0N0NXVBbDdZcGFLN2ViVXpvSzJDeGhCNHU2LzlmUVF1ZklNU0lZUGtzdDdNeHMwckJ2CnlLYVVzUUtCZ1FEYWU5bTkzR0NEdDh0dXpKdG9McHZCREJUbWlHTC9HWTJkcldrbkZiazFOU0RoWVNqTDFWNmoKUkloazZGS0p4aEZEV2JTdGlaVmtOcTQ0MnNxelhOR3UvTm9MZ1ltY2ZkL3dtanpKdEhQY1l4bWtHU2ZrQW9jRQorekFqLzIrZHIrVkorMW1oRDlaOFEram9hQXVOdFVNOWluOElDSnlEa0drclFrQVN0OWlSaXc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="))
//		return kconfig, err
//	})
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	// create the clientset
//	cli, err := dynamic.NewForConfig(config)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	//if err = json.Unmarshal([]byte(info), &deploy); err != nil {
//	//	fmt.Println(err.Error())
//	//}
//	obj := &unstructured.Unstructured{}
//	_, gvk, err := decUnstructured.Decode([]byte(info), nil, obj)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(obj)
//	namespace := "default"
//
//	groupList, err := discoveryClient.ServerPreferredResources()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	for _, list := range groupList {
//		if len(list.APIResources) == 0 {
//			continue
//		}
//		gv, err := schema.ParseGroupVersion(list.GroupVersion)
//		if err != nil {
//			fmt.Println(err.Error())
//			continue
//		}
//		for _, resource := range list.APIResources {
//			apiVersions[resource.Kind] = Resource{
//				Name:       resource.Name,
//				APIVersion: gv.String(),
//				Namespaced: resource.Namespaced,
//				Kind:       resource.Kind,
//			}
//		}
//	}
//	//apiVersions, err := ib.k8sCli.GetAPIVersion(ctx, req.ProjectId)
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//}
//	gvr := schema.GroupVersionResource{Group: gvk.Group, Version: gvk.Version, Resource: apiVersions[gvk.Kind].Name}
//	applyOption := metaV1.ApplyOptions{
//		FieldManager: "vine-dcp",
//		Force:        true,
//	}
//	_, err = cli.Resource(gvr).Namespace(namespace).Apply(ctx, obj.GetName(), obj, applyOption)
//
//}

func NewRuleBiz(
	logger log.Logger,
) *RuleBiz {
	return &RuleBiz{
		log: log.NewHelper(log.With(logger, "module", "biz/rule")),
	}
}
