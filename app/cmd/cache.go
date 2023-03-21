package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goapi/pkg/cache"
	"goapi/pkg/console"
)

var CacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CacheClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
}

var CacheForgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key, example: cache forget cache-key",
	Run:   runCacheForget,
}

// forget 命令的选项
var cacheKey string

func init() {
	// 注册 cache 命令的子命令
	CacheCmd.AddCommand(CacheClearCmd, CacheForgetCmd)

	// 设置 cache forget 命令的选项
	CacheForgetCmd.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	err := CacheForgetCmd.MarkFlagRequired("key")
	if err != nil {
		return
	}
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
