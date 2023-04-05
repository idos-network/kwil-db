package chainsync

import (
	"context"
	"fmt"
)

// Sync synchronizes the deposits and withdrawals up to the latest block.
func (c *chain) sync(ctx context.Context) error {
	// get the last synced height
	lastSynced, err := c.dao.GetHeight(ctx, int32(c.chainClient.ChainCode()))
	if err != nil {
		return fmt.Errorf("failed to get last synced height: %w", err)
	}

	// get the latest confirmed block
	latest, err := c.chainClient.GetLatestBlock(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest block: %w", err)
	}

	// split into chunks of n blocks
	chunks := splitBlocks(lastSynced, latest.Height, c.chunkSize)

	for _, chnk := range chunks {
		err = c.processChunk(ctx, chnk[0], chnk[1])
		if err != nil {
			return fmt.Errorf("failed to process chunk: %w", err)
		}
	}
	return nil
}