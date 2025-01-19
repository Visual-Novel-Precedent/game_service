package node

import (
	"github.com/rs/zerolog"
	node2 "visual_novel/internal/clients/node"
	"visual_novel/internal/models"
)

func GetNodesById(
	id int64,
	log *zerolog.Logger,
) (*[]models.Node, error) {
	var nodes *[]models.Node

	nodes, err := node2.GetNodesByChapterId(id, log)

	if err != nil {
		return nil, err
	}

	return nodes, nil
}
